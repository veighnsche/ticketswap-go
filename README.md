# Ticket Swap backend — TODO

1. Project foundation
- Keep the HTTP server and /health endpoint.
- Add environment variables for the server address, database URL, and auth secret.
- Connect to PostgreSQL.
- Make the server fail clearly if the database cannot connect.
- Add database migrations.
- Add a development command that runs the server with Air.
- 
2. Database schema
- users
  - id
  - email
  - password hash
  - display name
  - created at
- events
  - id
  - title
  - description
  - venue
  - city
  - starts at
  - image URL
  - created at
- tickets
  - id
  - event id
  - owner id
  - ticket type
  - original price
  - status: owned, listed, sold
- listings
  - id
  - ticket id
  - seller id
  - asking price
  - status: active, claimed, cancelled
  - created at
- orders
  - id
  - listing id
  - buyer id
  - seller id
  - price
  - status: completed
  - created at

3. Authentication
- Register a user.
- Log in with email and password.
- Hash passwords—never store plain passwords.
- Return an authentication token.
- Add middleware that identifies the logged-in user.
- Add GET /me.

4. Event API
- GET /events — list events.
- GET /events/:id — event details.
- POST /events — create an event.
- PATCH /events/:id — edit an event.
- DELETE /events/:id — delete an event.
- Decide later whether only an admin can create events.

5. Ticket API
- POST /events/:id/tickets — add a ticket to the logged-in user.
- GET /me/tickets — show the user’s tickets.
- GET /tickets/:id — ticket details.
- Prevent users from viewing or changing another user’s tickets.

6. Listing API
- POST /tickets/:id/listing — list an owned ticket for sale.
- Reject a ticket that is already listed or sold.
- GET /listings — browse active listings.
- Filter listings by event.
- GET /listings/:id — listing details.
- DELETE /listings/:id — seller cancels their own listing.

7. Claim a listing
- POST /listings/:id/claim — buyer claims an active listing.
- Use one database transaction.
- Ensure only one buyer can claim a listing.
- Mark the listing claimed.
- Transfer the ticket to the buyer.
- Create an order record.
- Return the updated ticket and order.

8. API quality
- Validate request bodies.
- Return consistent JSON errors.
- Use correct status codes.
- Add pagination to event and listing lists.
- Add search by event title or city.
- Add request logging.
- Keep secrets out of Git.

9. Frontend contract
- Write down each endpoint, request body, response body, and error case.
- Add CORS for the local frontend.
- Decide the frontend’s required screens:
  - event list
  - event detail
  - login/register
  - my tickets
  - create listing
  - listing checkout/claim
  - my purchases

10. Later, deliberately out of scope
- Payments and refunds.
- Real ticket verification.
- Email confirmation.
- Image uploads.
- Seller payouts.
- Admin dashboard.
- Fraud prevention.
- Notifications.

Start with: database connection → users → events → tickets → listings → claim transaction.
