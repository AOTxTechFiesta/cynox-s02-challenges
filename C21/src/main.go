package main

import (
	"crypto/rand"
	"database/sql"
	"encoding/hex"
	"errors"
	"fmt"
	"html/template"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"

	_ "github.com/mattn/go-sqlite3"
)

const (
	FLAG = "cyn0x{n0t_4ll_3ntr13s_ar3_l1st3d_but_4rt1cl3_4_w45_h1dd3n_n0t_d3l3t3d}"
)

// Article represents a blog article
type Article struct {
	ID      int
	Title   string
	Content string
	Owner   string
}

var (
	db           *sql.DB
	templates    *template.Template
	cookieName   = "owner_cookie"
	cookieMaxAge = 24 * 3600 * 30 // 30 days in seconds
)

// Initialize the database and create tables if they don't exist
func initDB() (*sql.DB, error) {
	database, err := sql.Open("sqlite3", "articles.db")
	if err != nil {
		return nil, err
	}
	return database, nil
}

func createArticleTable(db *sql.DB, tableName string) error {
	// Ensure the table name is valid and does not contain special characters
	if strings.ContainsAny(tableName, " ;'\"") {
		return fmt.Errorf("invalid table name: %s", tableName)
	}

	// Create articles table if it doesn't exist
	statement, err := db.Prepare(`
		CREATE TABLE IF NOT EXISTS ` + tableName + ` (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			title TEXT NOT NULL,
			content TEXT NOT NULL,
			owner TEXT NOT NULL
		)
	`)
	if err != nil {
		return err
	}
	_, err = statement.Exec()
	if err != nil {
		return err
	}
	return nil
}

// Generate a random alphanumeric string of the specified length
func generateRandomString(length int) (string, error) {
	if length < 1 {
		return "", errors.New("length must be at least 1")
	}

	// We need at least one byte to generate one letter, then (length-1)/2 for the rest
	bytes := make([]byte, (length+1)/2)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}

	hexStr := hex.EncodeToString(bytes)

	// Replace the first character with a random hex letter (a-f)
	letters := []rune{'a', 'b', 'c', 'd', 'e', 'f'}
	index, err := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))
	if err != nil {
		return "", err
	}

	hexRunes := []rune(hexStr)
	hexRunes[0] = letters[index.Int64()]

	return string(hexRunes[:length]), nil
}

// Get or create a cookie for the user
func getOrCreateCookie(w http.ResponseWriter, r *http.Request) (string, bool) {
	cookie, err := r.Cookie(cookieName)

	// If cookie doesn't exist, create a new one
	if err != nil || cookie.Value == "" {
		cookieValue, err := generateRandomString(16)
		if err != nil {
			log.Printf("Error generating cookie: %v", err)
			return "", false
		}

		// Set the new cookie
		newCookie := http.Cookie{
			Name:     cookieName,
			Value:    cookieValue,
			Path:     "/",
			MaxAge:   cookieMaxAge,
			HttpOnly: true,
			SameSite: http.SameSiteLaxMode,
		}
		http.SetCookie(w, &newCookie)

		err = createArticleTable(db, cookieValue)

		if err != nil {
			log.Printf("Error creating article table: %v", err)
		}

		// Insert some initial articles if they don't exist
		_, err = db.Exec("INSERT OR IGNORE INTO "+cookieValue+" (id, title, content, owner) VALUES (1, 'Understanding Web Security in 2025', 'Web security continues to evolve with new threats and solutions...', ?)", cookieValue)
		if err != nil {
			log.Printf("Error creating default article: %v", err)
		}

		_, err = db.Exec("INSERT OR IGNORE INTO "+cookieValue+" (id, title, content, owner) VALUES (2, 'The Evolution of Cyber Challenges', 'Cyber challenges have become increasingly sophisticated...', ?)", cookieValue)
		if err != nil {
			log.Printf("Error creating default article: %v", err)
		}

		_, err = db.Exec("INSERT OR IGNORE INTO "+cookieValue+" (id, title, content, owner) VALUES (3, 'Cryptography Trends for Next-Gen Security', 'Post-quantum cryptography is becoming increasingly important...', ?)", cookieValue)
		if err != nil {
			log.Printf("Error creating default article: %v", err)
		}

		_, err = db.Exec("INSERT OR IGNORE INTO "+cookieValue+" (id, title, content, owner) VALUES (4, 'FLAG', ?, ?)", FLAG, "cynox")
		if err != nil {
			log.Printf("Error creating default article: %v", err)
		}

		_, err = db.Exec("INSERT OR IGNORE INTO "+cookieValue+" (id, title, content, owner) VALUES (5, 'Reverse Engineering: Art and Science', 'Reverse engineering combines technical skills with creative problem-solving...', ?)", cookieValue)
		if err != nil {
			log.Printf("Error creating default article: %v", err)
		}

		_, err = db.Exec("INSERT OR IGNORE INTO "+cookieValue+" (id, title, content, owner) VALUES (6, 'Hidden Secrets in Plain Sight', 'Steganography allows hiding information in plain sight...', ?)", cookieValue)
		if err != nil {
			log.Printf("Error creating default article: %v", err)
		}

		return cookieValue, true
	}

	return cookie.Value, false
}

func main() {
	var err error
	db, err = initDB()
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}
	defer db.Close()

	// Parse all templates at startup
	templates = template.Must(template.ParseGlob("templates/*.html"))

	// Articles list endpoint
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		ownerCookie, _ := getOrCreateCookie(w, r)

		// Query articles that are either or owned by the current user
		// Validate that the ownerCookie does not contain special characters
		if strings.ContainsAny(ownerCookie, "'\"; ") {
			http.Error(w, "Invalid cookie value", http.StatusBadRequest)
			log.Printf("Invalid ownerCookie: %s", ownerCookie)
			return
		}

		rows, err := db.Query("SELECT id, title FROM "+ownerCookie+" WHERE owner = ?", ownerCookie)
		if err != nil {
			http.Error(w, "Database error", http.StatusInternalServerError)
			log.Printf("Database error: %v", err)
			return
		}
		defer rows.Close()

		type ArticlePreview struct {
			ID    int
			Title string
		}

		var articles []ArticlePreview
		for rows.Next() {
			var article ArticlePreview
			if err := rows.Scan(&article.ID, &article.Title); err != nil {
				log.Printf("Error scanning article: %v", err)
				continue
			}
			articles = append(articles, article)
		}

		data := map[string]interface{}{
			"Articles": articles,
		}

		templates.ExecuteTemplate(w, "index.html", data)
	})

	// New article creation endpoint
	http.HandleFunc("/articles/new", func(w http.ResponseWriter, r *http.Request) {
		// Get the user's cookie
		ownerCookie, _ := getOrCreateCookie(w, r)

		// Handle POST request - create new article
		if r.Method == http.MethodPost {
			// Parse form data
			err := r.ParseForm()
			if err != nil {
				http.Error(w, "Error parsing form", http.StatusBadRequest)
				return
			}

			title := r.FormValue("title")
			content := r.FormValue("content")

			// Validate input
			if title == "" || content == "" {
				data := map[string]interface{}{
					"ErrorMessage": "Title and content are required",
				}
				templates.ExecuteTemplate(w, "newarticle.html", data)
				return
			}

			// Insert new article into database
			result, err := db.Exec("INSERT INTO "+ownerCookie+" (title, content, owner) VALUES (?, ?, ?)",
				title, content, ownerCookie)

			if err != nil {
				log.Printf("Error creating article: %v", err)
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}

			// Get the ID of the newly inserted article
			articleID, err := result.LastInsertId()
			if err != nil {
				log.Printf("Error getting last insert ID: %v", err)
				http.Error(w, "Database error", http.StatusInternalServerError)
				return
			}

			// Redirect to the new article
			http.Redirect(w, r, "/article/"+strconv.FormatInt(articleID, 10), http.StatusSeeOther)
			return
		}

		// Handle GET request - show creation form
		data := map[string]interface{}{
			"ErrorMessage": "",
		}

		templates.ExecuteTemplate(w, "newarticle.html", data)
	})

	// View article endpoint
	http.HandleFunc("/article/", func(w http.ResponseWriter, r *http.Request) {
		parts := strings.Split(r.URL.Path, "/")

		// Handle view article path: /article/{id}
		if len(parts) == 3 && parts[1] == "article" {
			articleID, err := strconv.Atoi(parts[2])
			if err != nil {
				http.NotFound(w, r)
				return
			}

			// Get the user's cookie
			ownerCookie, _ := getOrCreateCookie(w, r)

			// Validate that the ownerCookie does not contain special characters
			if strings.ContainsAny(ownerCookie, "'\"; ") {
				http.Error(w, "Invalid cookie value", http.StatusBadRequest)
				log.Printf("Invalid ownerCookie: %s", ownerCookie)
				return
			}

			// Fetch the article from the database
			var article Article
			err = db.QueryRow("SELECT id, title, content, owner FROM "+ownerCookie+" WHERE id = ?", articleID).
				Scan(&article.ID, &article.Title, &article.Content, &article.Owner)

			if err != nil {
				if err == sql.ErrNoRows {
					http.NotFound(w, r)
				} else {
					http.Error(w, "Database error", http.StatusInternalServerError)
					log.Printf("Database error: %v", err)
				}
				return
			}

			// Check if the user is allowed to view this article
			data := map[string]interface{}{
				"ID":        article.ID,
				"Title":     article.Title,
				"Content":   article.Content,
				"Forbidden": false,
			}

			if article.Owner != ownerCookie {
				data["Forbidden"] = true
				data["Title"] = "Access Forbidden"
				data["Content"] = ""
			}

			templates.ExecuteTemplate(w, "article.html", data)
			return
		}

		// Handle article edit paths: /article/edit/{id}
		if len(parts) == 4 && parts[1] == "article" && parts[2] == "edit" {
			articleID, err := strconv.Atoi(parts[3])
			if err != nil {
				http.NotFound(w, r)
				return
			}

			// Get the user's cookie
			ownerCookie, _ := getOrCreateCookie(w, r)

			// Validate that the ownerCookie does not contain special characters
			if strings.ContainsAny(ownerCookie, "'\"; ") {
				http.Error(w, "Invalid cookie value", http.StatusBadRequest)
				log.Printf("Invalid ownerCookie: %s", ownerCookie)
				return
			}

			// Fetch the article from the database
			var article Article
			err = db.QueryRow("SELECT id, title, content, owner FROM "+ownerCookie+" WHERE id = ?", articleID).
				Scan(&article.ID, &article.Title, &article.Content, &article.Owner)

			if err != nil {
				if err == sql.ErrNoRows {
					http.NotFound(w, r)
				} else {
					http.Error(w, "Database error", http.StatusInternalServerError)
					log.Printf("Database error: %v", err)
				}
				return
			}

			// Handle POST request - update article
			if r.Method == http.MethodPost {
				// For POST, check if article belongs to user
				if article.Owner != ownerCookie {
					http.Error(w, "Forbidden: You don't have permission to edit this article", http.StatusForbidden)
					return
				}

				// Parse form data
				err := r.ParseForm()
				if err != nil {
					http.Error(w, "Error parsing form", http.StatusBadRequest)
					return
				}

				title := r.FormValue("title")
				content := r.FormValue("content")

				// Validate that the ownerCookie does not contain special characters
				if strings.ContainsAny(ownerCookie, "'\"; ") {
					http.Error(w, "Invalid cookie value", http.StatusBadRequest)
					log.Printf("Invalid ownerCookie: %s", ownerCookie)
					return
				}

				// Update article in database
				_, err = db.Exec("UPDATE "+ownerCookie+" SET title = ?, content = ? WHERE id = ?",
					title, content, articleID)

				if err != nil {
					log.Printf("Error updating article: %v", err)
					http.Error(w, "Database error", http.StatusInternalServerError)
					return
				}

				// Redirect to the article view page
				http.Redirect(w, r, "/article/"+parts[3], http.StatusSeeOther)
				return
			}

			// Handle GET request - show edit form
			// For the GET request, we don't check permissions,
			// but we will prepare the editing form
			data := map[string]interface{}{
				"ID":           article.ID,
				"Title":        article.Title,
				"Content":      article.Content,
				"ErrorMessage": "",
			}

			templates.ExecuteTemplate(w, "editor.html", data)
			return
		}

		// If path doesn't match any expected format
		http.NotFound(w, r)
	})

	log.Println("Server starting on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
