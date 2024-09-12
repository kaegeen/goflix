document.addEventListener('DOMContentLoaded', function() {
    // Load movies on page load
    loadMovies();

    // Login form submission
    document.getElementById('login-form').addEventListener('submit', function(e) {
        e.preventDefault();
        const username = document.getElementById('username').value;
        const password = document.getElementById('password').value;

        // Perform the login request
        fetch('/login', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({ username, password })
        })
        .then(response => response.json())
        .then(data => {
            if (data.token) {
                localStorage.setItem('token', data.token);
                document.getElementById('login-message').innerText = 'Login successful!';
            } else {
                document.getElementById('login-message').innerText = 'Login failed!';
            }
        });
    });

    // Add movie form submission
    document.getElementById('add-movie-form').addEventListener('submit', function(e) {
        e.preventDefault();
        const title = document.getElementById('title').value;
        const releaseDate = document.getElementById('release-date').value;
        const duration = document.getElementById('duration').value;
        const trailerUrl = document.getElementById('trailer-url').value;

        const token = localStorage.getItem('token');
        if (!token) {
            document.getElementById('add-movie-message').innerText = 'Please log in first!';
            return;
        }

        fetch('/api/movies', {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer ' + token
            },
            body: JSON.stringify({ title, release_date: releaseDate, duration, trailer_url: trailerUrl })
        })
        .then(response => response.json())
        .then(data => {
            if (data.id) {
                document.getElementById('add-movie-message').innerText = 'Movie added successfully!';
                loadMovies();
            } else {
                document.getElementById('add-movie-message').innerText = 'Failed to add movie!';
            }
        });
    });
});

// Function to load movies
function loadMovies() {
    fetch('/api/movies', {
        headers: {
            'Content-Type': 'application/json'
        }
    })
    .then(response => response.json())
    .then(data => {
        const moviesContainer = document.getElementById('movies');
        moviesContainer.innerHTML = '';

        data.forEach(movie => {
            const movieItem = document.createElement('div');
            movieItem.classList.add('movie-item');

            movieItem.innerHTML = `
                <h3>${movie.title}</h3>
                <p>Release Date: ${movie.release_date}</p>
                <p>Duration: ${movie.duration} minutes</p>
                <a href="${movie.trailer_url}" target="_blank">Watch Trailer</a>
            `;

            moviesContainer.appendChild(movieItem);
        });
    });
}
