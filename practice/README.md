# Practice: OnePizza

[Link to the original page](https://codingcompetitions.withgoogle.com/hashcode/round/00000000008f5ca9/00000000008f6f33).

You are opening a small pizzeria. In fact, your pizzeria is so small that you decided to offer only one type of pizza. Now you need to decide what ingredients to include (peppers? tomatoes? both?).

Everyone has their own pizza preferences. Each of your potential clients has some ingredients they like, and maybe some ingredients they dislike. Each client will come to your pizzeria if both conditions are true:

    1. all the ingredients they like are on the pizza, and
    2. none of the ingredients they dislike are on the pizza

Each client is OK with additional ingredients they neither like or dislike being present on the pizza. Your task is to choose which ingredients to put on your only pizza type, to maximize the number of clients that will visit your pizzeria.
Input

    The first line contains one integer 1≤C≤105 - the number of potential clients.
    The following 2×C lines describe the clients’ preferences in the following format:
        First line contains integer 1≤L≤5, followed by L names of ingredients a client likes, delimited by spaces.
        Second line contains integer 0≤D≤5, followed by D names of ingredients a client dislikes, delimited by spaces.

Each ingredient name consists of between 1 and 15 ASCII characters. Each character is one of the lowercase letters (a-z) or a digit (0-9).
