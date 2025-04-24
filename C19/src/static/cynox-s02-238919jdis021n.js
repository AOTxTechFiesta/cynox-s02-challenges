function getFlag() {
  const endpoint = "L2dldC1mbGFnLWU4MmgxM2g1Mm4yMw==";

  fetch(atob(endpoint), {
    method: "POST",
  })
    .then((response) => response.text())
    .then((data) => {
      console.log("Flag:", data);
      alert("Flag: " + data);
    })
    .catch((error) => {
      console.error("Error:", error);
      alert("Error retrieving flag. Please try again.");
    });
}