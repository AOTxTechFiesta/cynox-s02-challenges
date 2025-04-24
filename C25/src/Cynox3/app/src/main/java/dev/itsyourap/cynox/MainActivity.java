package dev.itsyourap.cynox;

import androidx.activity.EdgeToEdge;
import androidx.appcompat.app.AppCompatActivity;
import androidx.core.graphics.Insets;
import androidx.core.view.ViewCompat;
import androidx.core.view.WindowInsetsCompat;

import android.os.Bundle;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;

public class MainActivity extends AppCompatActivity {
    private String uO9rPtdf5aB3rGnDT21u7Sq0ASlFUN = null;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        EdgeToEdge.enable(this);
        setContentView(R.layout.activity_main);
        ViewCompat.setOnApplyWindowInsetsListener(findViewById(R.id.main), (v, insets) -> {
            Insets systemBars = insets.getInsets(WindowInsetsCompat.Type.systemBars());
            v.setPadding(systemBars.left, systemBars.top, systemBars.right, systemBars.bottom);
            return insets;
        });

        EditText flagEditText = findViewById(R.id.flagEditText);
        Button submitButton = findViewById(R.id.submitButton);
        TextView messageTextView = findViewById(R.id.messageTextView);

        submitButton.setOnClickListener(v -> {
            String flag = flagEditText.getText().toString();
            if (uO9rPtdf5aB3rGnDT21u7Sq0ASlFUN == null){
                uO9rPtdf5aB3rGnDT21u7Sq0ASlFUN = new C28421().Ey1vO4MRjn379HD0RTYrCfHQQZ2dK4();
            }

            if (flag.equals(uO9rPtdf5aB3rGnDT21u7Sq0ASlFUN)) {
                messageTextView.setText(getResources().getText(R.string.correct_flag));
            } else {
                messageTextView.setText(getResources().getText(R.string.incorrect_flag));
            }
        });
    }
}