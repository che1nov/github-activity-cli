# GitHub Activity CLI
https://roadmap.sh/projects/github-user-activity

This project is a simple Command Line Interface (CLI) application that fetches and displays recent GitHub activity for a specified user. This project will help you practice your programming skills, including working with APIs, handling JSON data, and creating a simple CLI application.

## Features

- Fetch recent GitHub activity for a specified user.
- Display the fetched activity in the terminal.
- Handle errors such as invalid usernames or API failures gracefully.

## Requirements

- The application should be run from the command line, accepting a GitHub username as an argument.
- Fetch recent activity for the specified GitHub user using the GitHub API.
- Display the fetched activity in the terminal.

## Usage

To run the application, use the following command:

```
github-activity <username>
```

Replace `<username>` with the GitHub username you want to fetch activity for.

Example:

```
github-activity kamranahmedse
```

## GitHub API Endpoint

To fetch the user's recent activity, the application uses the following GitHub API endpoint:

```
https://api.github.com/users/<username>/events
```

Replace `<username>` with the actual GitHub username.

Example:

```
https://api.github.com/users/kamranahmedse/events
```

## Output

The application will display the fetched activity in the terminal. Example output:

```
- Pushed 3 commits to kamranahmedse/developer-roadmap
- Opened a new issue in kamranahmedse/developer-roadmap
- Starred kamranahmedse/developer-roadmap
...
```

## Error Handling

The application will handle errors gracefully, including:

- Invalid GitHub usernames.
- API failures.

## Advanced Features

For a more advanced version of this project, you can consider adding features such as:

- Filtering activity by event type.
- Displaying activity in a more structured format.
- Caching fetched data for improved performance.
- Exploring other GitHub API endpoints to fetch additional information about the user or their repositories.

## Development

To develop this project, choose your preferred programming language. Do not use external libraries or frameworks for fetching GitHub activity.

## License

This project is licensed under the MIT License.

## Acknowledgments

- GitHub API Documentation: [GitHub API](https://docs.github.com/en/rest)