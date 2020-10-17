# API Specification
There are 3 paths avaliable in our API.

* ##/r
   
   Allows method POST - creates new shortenedURL (requires key)

* ##/r/{urlBase58ID}
   Allows method GET - accesses shortened URL and redirects to original URL (public)
   
   Allows method DELETE - deletes shortened URL (requires key) - This will also clear the stats for that URL
 
* ##/s/{urlBase58ID}?time={timeframe}
   Allows method GET - Retrieves url access statistics (requires key)
   