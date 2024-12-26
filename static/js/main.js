let imageInfo = {};
// Function to load a random cat image from the Beego backend
const loadImage = async () => {
  const loader = document.getElementById("loader");
  const img = document.getElementById("cat-image");
  loader.style.display = "block";
  img.style.display = "none";
  const response = await fetch("/cat/get-image");
  const data = await response.json();
  imageInfo = data;

  if (data && data.length > 0) {
    img.src = data[0].url;
    img.style.display = "block";
    img.alt = "Random Cat";
    loader.style.display = "none";
  } else {
    img.style.display = "block";
    loader.style.display = "none";
    container.innerHTML = "<p>Failed to fetch cat image.</p>";
  }
};

// Function to handle the "like" button click
const likeImage = () => {
  loadImage();
};

// Function to handle the "dislike" button click
const dislikeImage = () => {
  loadImage();
};

const favoriteCat = () => {
  loadImage();
  const imageId = imageInfo[0].id; 
  const subId = "user-1234"; 
  console.log("imageId",imageId)

  fetch("/api/favorites", {
    method: "POST",
    headers: { "Content-Type": "application/json" },
    body: JSON.stringify({ image_id: imageId, sub_id: subId }),
  })
    .then((res) => {
      if (!res.ok) {
        throw new Error(`HTTP error! Status: ${res.status}`);
      }
      return res.json();
    })
    .then((data) => {
      console.log(data)
       // Handle the response data
    })
    .catch((error) => {
      console.error("Error:", error); 
    });
};

// Add click event listeners to the like and dislike buttons
document.querySelector(".fa-thumbs-up").addEventListener("click", likeImage);
document
  .querySelector(".fa-thumbs-down")
  .addEventListener("click", dislikeImage);

// document.querySelector(".fa-heart").addEventListener("click", favoriteCat);

// Add click handlers for navigation tabs
document.querySelectorAll(".nav-item").forEach((tab) => {
  tab.addEventListener("click", () => {
    document
      .querySelectorAll(".nav-item")
      .forEach((t) => t.classList.remove("active"));
    tab.classList.add("active");
  });
});

// Add click handlers for interaction buttons
document.querySelectorAll(".icon-button").forEach((button) => {
  button.addEventListener("click", () => {
    button.style.transform = "scale(1.2)";
    setTimeout(() => {
      button.style.transform = "scale(1)";
    }, 200);
  });
});

// Load favorites on page load
window.onload = loadFavorites;
