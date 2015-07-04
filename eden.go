package main

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/fortytw2/eden/api"
	"github.com/fortytw2/eden/datastore"
	"github.com/fortytw2/eden/datastore/pgsql"
	"github.com/fortytw2/eden/datastore/pgsql/queries"
	"github.com/fortytw2/eden/datastore/redis"
	"github.com/fortytw2/eden/web"
	"github.com/julienschmidt/httprouter"

	// autoload ENV from .env
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	log.Println("eden: initializing")
	log.Printf("eden: loaded %d sql queries\n", len(queries.All()))

	data := constructDatastore()

	router := httprouter.New()
	router.GET("/", web.Homepage)

	router.GET("/b/", api.GetBoards(data))
	router.POST("/b/", api.CreateBoard(data))

	router.GET("/b/:board/", api.GetBoardPosts(data))
	router.GET("/b/:board/post/:id", api.GetPost(data))
	router.POST("/b/:board/", api.CreatePost(data))
	router.POST("/b/:board/post/:id", api.CreateComment(data))

	router.POST("/u/", api.NewUser(data))
	// router.GET("/u/:username")

	log.Println("eden: now listening on port", os.Getenv("PORT"))
	err := http.ListenAndServe(os.Getenv("PORT"), httpLogger(router))
	if err != nil {
		panic(err)
	}
}

// cleanly log all HTTP requests
func httpLogger(router http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		startTime := time.Now()
		router.ServeHTTP(w, req)
		finishTime := time.Now()
		elapsedTime := finishTime.Sub(startTime)
		log.Println("eden:", req.Method, req.URL, elapsedTime)
	})
}

// build our production datastore
func constructDatastore() *datastore.Datastore {
	db, err := pgsql.NewDBHandle()
	if err != nil {
		log.Fatalf("eden: fatal error getting DB handle, %s\n", err)
	}
	r := redis.GetRedisPool()

	return &datastore.Datastore{
		UserService:    pgsql.NewUserService(db),
		BoardService:   pgsql.NewBoardService(db),
		PostService:    pgsql.NewPostService(db),
		CommentService: pgsql.NewCommentService(db),
		VoteService:    redis.NewVoteService(r),
	}
}
