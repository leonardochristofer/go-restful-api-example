# Cakes - Go RESTful API Example

### How to run

1. Install Docker https://docs.docker.com/get-docker/

2. Use make run / docker-compose up --build to run the app

3. Use make stop / docker-compose down to stop the app

4. Use make fmt to tidy up code style

5. make mock-gen to generate mockery file

6. make test to view unit testing

### How to contribute

If you'd like to contribute to this repository, follow these steps:

1. Create a new branch with a descriptive name for your changes.

2. Make your proposed changes and improvements in your branch.

3. Test your changes to ensure they are accurate and don't introduce any issues.

4. Commit your changes and push them to your branch (commit message guide: https://www.conventionalcommits.org/en/v1.0.0/).

5. Create a pull request from your branch to the main branch.

### API Example

1. curl 'http://localhost:3200/api/cakes'

2. curl 'http://localhost:3200/api/cakes/2'

3. curl 'http://localhost:3200/api/cakes' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Velvet Dream",
    "description": "A luscious red velvet cake, with layers of moist chocolate sponge, topped with smooth cream cheese frosting and garnished with chocolate shavings.",
    "rating": 10,
    "image": "https://example.com/111"
}'

4. curl --request PUT 'http://localhost:3200/api/cakes/1' \
--header 'Content-Type: application/json' \
--data '{
    "title": "Lemon Zest Delight",
    "description": "Indulge in the refreshing flavors of our Lemon Zest Delight cake.",
    "rating": 7,
    "image": "https://example.com/333"
}'

5. curl --request DELETE 'http://localhost:3200/api/cakes/2'