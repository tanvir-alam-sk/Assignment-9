const favoriteCatList = [];

// Function to load a random cat image from the Beego backend
const loadImage = async () => {
    const loader = document.getElementById('loader');
    const img = document.getElementById('cat-image');
    loader.style.display = 'block';
    img.style.display = 'none';
    const response = await fetch('/cat/get-image');
    const data = await response.json();


    if (data && data.length > 0) {
        img.src = data[0].url;
        img.style.display = 'block';
        img.alt = "Random Cat";
        loader.style.display = 'none';

    } else {
        img.style.display = 'block';
        loader.style.display = 'none';
        container.innerHTML = '<p>Failed to fetch cat image.</p>';
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

// Add click event listeners to the like and dislike buttons
document.querySelector(".fa-thumbs-up").addEventListener("click", likeImage);
document.querySelector(".fa-thumbs-down").addEventListener("click", dislikeImage);



// Retrieve data from Local Storage on page load (if needed)
const loadFavorites = () => {
    const savedFavorites = JSON.parse(localStorage.getItem("favCat")) || [];
};

// Add click handlers for navigation tabs
document.querySelectorAll('.nav-item').forEach(tab => {
    tab.addEventListener('click', () => {
        document.querySelectorAll('.nav-item').forEach(t => t.classList.remove('active'));
        tab.classList.add('active');
    });
});

// Add click handlers for interaction buttons
document.querySelectorAll('.icon-button').forEach(button => {
    button.addEventListener('click', () => {
        button.style.transform = 'scale(1.2)';
        setTimeout(() => {
            button.style.transform = 'scale(1)';
        }, 200);
    });
});

// Load favorites on page load
window.onload = loadFavorites;