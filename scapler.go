package scapler

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/joho/godotenv"
)

const version = "1.0.0"

type Scapler struct {
	AppName  string
	Debug    bool
	Version  string
	ErrorLog *log.Logger
	InfoLog  *log.Logger
	RootPath string
	Routes   *chi.Mux
	config   config
}

type config struct {
	port     string
	renderer string
}

func (s *Scapler) New(rootPath string) error {
	pathConfig := initPaths{
		rootPath: rootPath,
		folderNames: []string{
			"handlers", "migrations", "views", "data", "public", "logs", "middleware",
		},
	}

	err := s.Init(pathConfig)
	if err != nil {
		return err
	}

	err = s.checkDotEnv(rootPath)
	if err != nil {
		return err
	}

	err = godotenv.Load(rootPath + "/.env")
	if err != nil {
		return err
	}

	infoLog, errorLog := s.startLoggers()
	s.InfoLog = infoLog
	s.ErrorLog = errorLog
	s.Debug, _ = strconv.ParseBool(os.Getenv("DEBUG"))
	s.Version = version
	s.RootPath = rootPath
	s.Routes = s.routes().(*chi.Mux)

	s.config = config{
		port:     os.Getenv("PORT"),
		renderer: os.Getenv("RENDERER"),
	}

	return nil
}

func (s *Scapler) Init(p initPaths) error {
	root := p.rootPath
	for _, path := range p.folderNames {
		err := s.CreateDirIfNotExists(root + "/" + path)
		if err != nil {
			return err
		}
	}
	return nil
}

func (s *Scapler) ListenAndServe() {
	srv := &http.Server{
		Addr:         fmt.Sprintf(":%s", os.Getenv("PORT")),
		ErrorLog:     s.ErrorLog,
		Handler:      s.routes(),
		IdleTimeout:  30 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 600 * time.Second,
	}

	s.InfoLog.Printf("Listening on port %s", os.Getenv("PORT"))
	err := srv.ListenAndServe()
	s.ErrorLog.Fatal(err)
}

func (s *Scapler) checkDotEnv(path string) error {
	err := s.CreateFileIfNotExists(fmt.Sprintf("%s/.env", path))
	if err != nil {
		return err
	}
	return nil
}

func (s *Scapler) startLoggers() (*log.Logger, *log.Logger) {
	var infoLog *log.Logger
	var errorLog *log.Logger

	infoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)

	return infoLog, errorLog
}
