package main

import (
    "encoding/json"
    "html/template"
    "io/ioutil"
    "net/http"
    "strconv"
    "strings"
    "sync"
)

type Todo struct {
    Title    string
    Done     bool
    Deadline string
    Tag      string
}

var (
    todos []Todo
    mu    sync.Mutex
    tmpl  = template.Must(template.ParseFiles("templates/index.html"))
)

func main() {
    loadTodos()

    http.HandleFunc("/", showTodos)
    http.HandleFunc("/add", addTodo)
    http.HandleFunc("/toggle", toggleDone)
    http.HandleFunc("/delete", deleteTodo)
    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

    http.ListenAndServe(":8080", nil)
}

func showTodos(w http.ResponseWriter, r *http.Request) {
    mu.Lock()
    defer mu.Unlock()

    query := strings.ToLower(r.URL.Query().Get("q"))
    var filtered []Todo
    for _, todo := range todos {
        if query == "" || strings.Contains(strings.ToLower(todo.Title), query) || strings.Contains(strings.ToLower(todo.Tag), query) {
            filtered = append(filtered, todo)
        }
    }

    tmpl.Execute(w, filtered)
}

func addTodo(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    title := r.FormValue("title")
    deadline := r.FormValue("deadline")
    tag := r.FormValue("tag")

    if title != "" {
        mu.Lock()
        todos = append(todos, Todo{Title: title, Deadline: deadline, Tag: tag})
        saveTodos()
        mu.Unlock()
    }

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func toggleDone(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    indexStr := r.FormValue("index")
    index, err := strconv.Atoi(indexStr)
    if err != nil || index < 0 || index >= len(todos) {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    mu.Lock()
    todos[index].Done = !todos[index].Done
    saveTodos()
    mu.Unlock()

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    indexStr := r.FormValue("index")
    index, err := strconv.Atoi(indexStr)
    if err != nil || index < 0 || index >= len(todos) {
        http.Redirect(w, r, "/", http.StatusSeeOther)
        return
    }

    mu.Lock()
    todos = append(todos[:index], todos[index+1:]...)
    saveTodos()
    mu.Unlock()

    http.Redirect(w, r, "/", http.StatusSeeOther)
}

func saveTodos() {
    data, err := json.MarshalIndent(todos, "", "  ")
    if err == nil {
        ioutil.WriteFile("todos.json", data, 0644)
    }
}

func loadTodos() {
    data, err := ioutil.ReadFile("todos.json")
    if err == nil {
        json.Unmarshal(data, &todos)
    }
}
