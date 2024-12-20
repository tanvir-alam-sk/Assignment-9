<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>The Cat API</title>
    <link href="https://cdnjs.cloudflare.com/ajax/libs/font-awesome/6.5.0/css/all.min.css" rel="stylesheet">
    <style>
        * {
            margin: 0;
            padding: 0;
            box-sizing: border-box;
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, Oxygen, Ubuntu, Cantarell, sans-serif;
        }

        body {
            background-color: #fff;
            min-height: 100vh;
            display: flex;
            flex-direction: column;
        }

        .container {
            max-width: 1200px;
            margin: 0 auto;
            padding: 2rem;
            display: grid;
            grid-template-columns: 1fr 1fr;
            gap: 2rem;
            align-items: start;
        }

        .left-section {
            padding-top: 2rem;
        }

        .title-highlight {
            background-color: #ffd7d7;
            padding: 0 0.5rem;
            display: inline-block;
        }

        h1 {
            font-size: 4rem;
            font-weight: 900;
            line-height: 1.1;
            margin-bottom: 1rem;
        }

        .subtitle {
            font-size: 2.5rem;
            font-weight: 900;
            margin-bottom: 2rem;
        }

        .description {
            color: #666;
            margin-bottom: 1rem;
        }

        .stats {
            color: #666;
            margin-bottom: 2rem;
        }

        .buttons {
            display: flex;
            gap: 1rem;
        }

        .btn {
            padding: 0.75rem 1.5rem;
            border-radius: 4px;
            font-weight: 600;
            cursor: pointer;
            transition: all 0.2s;
        }

        .btn-primary {
            background-color: #000;
            color: white;
            border: none;
        }

        .btn-secondary {
            background-color: white;
            color: #000;
            border: 1px solid #000;
        }

        .btn:hover {
            opacity: 0.9;
            transform: translateY(-1px);
        }

        .right-section {
            background-color: #fff;
            border-radius: 12px;
            box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
            overflow: hidden;
        }

        .nav-tabs {
            display: flex;
            padding: 1rem;
            border-bottom: 1px solid #eee;
        }
        a{
            text-decoration: none;
        }
        .nav-item {
            display: flex;
            flex-direction: column;
            align-items: center;
            gap: 0.5rem;
            margin-right: 2rem;
            color: #666;
            cursor: pointer;
        }

        .nav-item  .fa-heart{
            margin: 0;
        }

        .nav-item.active {
            color: #de7d0f;

        }

        .nav-item.active span{
            font-weight: 600;
        }        

        .cat-image {
            width: 600px;
            height: 400px;
            object-fit: cover;
        }

        .interaction-buttons {
            display: flex;
            justify-content: space-between;
            padding: 1rem;
        }

        .fa-regular{
            background: none;
            border: none;
            cursor: pointer;
            font-size: 1.5rem;
            color: #666;
            margin: 5px;
        }
        .fa-up-down ,.fa-magnifying-glass,.fa-heart{
            font-size: 20px;
        }

        .fa-regular:hover {
            color: #ffb340;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="left-section">
            <h1>
                <span class="title-highlight">The Cat API</span><br>
                Cats as a service.
            </h1>
            <div class="subtitle">Because everyday is a Caturday.</div>
            <p class="description">An API all about cat.</p>
            <p class="stats">60k+ Images. Breeds. Facts.</p>
            <div class="buttons">
                <button class="btn btn-primary">GET YOUR API KEY</button>
                <button class="btn btn-secondary">READ OUR GUIDES</button>
            </div>
        </div>
        <div class="right-section">
            <div class="nav-tabs">
                <a href="/home">
                <div class="nav-item active">
                        <i class="fa-solid fa-up-down"></i>
                        <span>Voting</span>
                    </div>
                </a>
                <div class="nav-item">
                    <i class="fa-solid fa-magnifying-glass"></i>
                    <span>Breeds</span>
                </div>
                <div class="nav-item">
                    <i class="fa-regular fa-heart"></i>
                    <span>Favs</span>
                </div>
            </div>
            <img src="https://cdn2.thecatapi.com/images/MTkzMzUzNg.jpg" alt="Cat" class="cat-image">
            <div class="interaction-buttons">
                <i class="fa-regular fa-heart" onclick="favoriteCat()"></i>
                <div class="fa-regular-btn">
                    <i class="fa-regular fa-thumbs-up" id="fa-thumbs-up"></i>
                    <i class="fa-regular fa-thumbs-down"></i>
                </div>
            </div>
        </div>
    </div>

    <script>

const favoriteCatList = [];

        // Function to load a random cat image from the Beego backend
        const loadImage = async () => {
            try {
                const catImageElement = document.querySelector(".cat-image"); 
                const response = await fetch("/cat/image");
                console.log(response)
                if (!response.ok) throw new Error("Failed to fetch image");
                const data = await response.json();
                catImageElement.setAttribute("src", data.url); 
            } catch (error) {
                console.log("Error loading image:", error);
            }
        };

        // Function to handle the "like" button click
        const likeImage = () => {
            console.log("Liked the image!"); // You can expand this to save the liked image
            loadImage(); // Load a new image
        };

        // Function to handle the "dislike" button click
        const dislikeImage = () => {
            console.log("Disliked the image!"); // You can expand this to save disliked images
            loadImage(); // Load a new image
        };

        // Add click event listeners to the like and dislike buttons
        document.querySelector(".fa-thumbs-up").addEventListener("click", likeImage);
        document.querySelector(".fa-thumbs-down").addEventListener("click", dislikeImage);

        // Retrieve data from Local Storage on page load (if needed)
        const loadFavorites = () => {
            const savedFavorites = JSON.parse(localStorage.getItem("favCat")) || [];
            console.log(savedFavorites); // You can use this to display the list of favorites if desired
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
    </script>
</body>
</html>