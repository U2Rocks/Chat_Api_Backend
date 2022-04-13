# Chat Messaging Api

## Summary

- This is a backend api that stores chat messages and allows messages to be retrieved in multiple ways

## Build

- This backend uses **Fiber** to setup the api routes and **Gorm** to handle interacting with the sqlite database

## Routes

- User Routes: /users -> /users/update -> /users/:id
- Chatroom Routes: /chatrooms -> /chatrooms/newtitle -> /chatrooms/newdesc -> /chatrooms/:id
- Message Routes: /messages -> /messages/chatroom -> /messages/textsearch -> /messages/:id

## Final Comments and Notes

- The admin panel at the base route for the api is purely decorative and does not have javascript hooked up to the page
- The extra style sheet was created to try and solve css issues
