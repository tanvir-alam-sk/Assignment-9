const favoriteImage = () => {
  const favorite = document.getElementById("image-section");
  favorite.innerHTML = "";
  fetch("/api/favorites")
    .then((response) => response.json())
    .then((data) => {
      data.data.forEach((datum) => {
        const div = document.createElement("div");
        div.classList.add("image-card");
        const img = document.createElement("img");
        img.src = datum.image.url;
        img.alt = "fav image";
        div.appendChild(img);
        // console.log(datum)
        const icon = document.createElement("i");
        icon.className = "fa-solid fa-trash";
        icon.addEventListener("click", () => deleteFav(datum.id));
        div.appendChild(icon);
        favorite.appendChild(div);
      });
    });
};

favoriteImage();

const deleteFav = (favouriteId) => {
  const requestOptions = {
    method: "DELETE",
    headers: {
      "Content-Type": "application/json",
    },
  };

  fetch(`/favourite/${favouriteId}`, requestOptions)
    .then((response) => {
      if (response.ok) {
        favoriteImage()
        console.log("Favourite deleted successfully");
      } else {
        console.error("Failed to delete favourite");
      }
    })
    .catch((error) => {
      console.error("Error:", error);
    });
};

const gridButton = document.getElementById("grid-view-btn");
const listButton = document.getElementById("list-view-btn");
const imageSection = document.getElementById("image-section");

// Function to toggle active button state
function setActiveButton(activeButton, inactiveButton) {
  activeButton.classList.add("active"); // Add active class to clicked button
  inactiveButton.classList.remove("active"); // Remove active class from the other button
}

// Switch to grid view
gridButton.addEventListener("click", () => {
  imageSection.classList.remove("list-view"); // Remove list view class
  setActiveButton(gridButton, listButton); // Update button states
});

// Switch to list view
listButton.addEventListener("click", () => {
  imageSection.classList.add("list-view"); // Add list view class
  setActiveButton(listButton, gridButton); // Update button states
});
