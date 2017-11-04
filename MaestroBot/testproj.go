package main

import (
    "crypto/md5"
    "encoding/hex"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "net/http/httptest"
    "strconv"
    "strings"
    "time"
    "os"
    //"encoding/json"
    "io/ioutil"
    //"log"
    //"net/http"
    //"bytes"



    cors "github.com/heppu/simple-cors"
)





var erro Err

type Feed3 struct{
    Tracks struct{
            Track [] struct{
                Name string
            }
        

    }
}


type Feed2 struct{
    Album struct{
        Artist string
        Tracks struct{
            Track [] struct{
                Name string
            }
        }

    }
}

type Feed1 struct{
    TopTracks struct{
        Track [] struct{
            Name string
            Playcount string
        }
    }

}

type track struct{
    Name string
    Playcount string
}

type album struct{
            Name string `json:"Name"`
            Playcount int `json:"Playcount"`
            Artist string `json:"Artist"`
        
} 

type Err struct{
    Error int `json:"Error"`
    Message string `json:"Message"`
    Links [] string `json:"Links"`
}

//import _ "github.com/joho/godotenv/autoload"
type Artist struct {
    Name     string `json:"Name"`
    //Author    string
    Url       string `json:"Url"`
    //Permalink string
}

// the feed is the full JSON data structure
// this sets up the array of Entry types (defined above)
type Feed struct {
    TopAlbums struct {
        Album []struct {
            Name string `json:"Name"`
            Playcount int `json:"Playcount"`
            Artist Artist
        }
    }
}

var (
    // WelcomeMessage A constant to hold the welcome message
    WelcomeMessage = "Welcome, what do you want to order?"

    // sessions = {
    //   "uuid1" = Session{...},
    //   ...
    // }
    sessions = map[string]Session{}

    processor = sampleProcessor
)

type (
    // Session Holds info about a session
    Session map[string]interface{}

    // JSON Holds a JSON object
    JSON map[string]interface{}

    // Processor Alias for Process func
    Processor func(session Session, message string) (string, string,string, string, string, error)
)


func fetchchart() (string,string, string, string, string, error) {

var entries Feed3
 //formattedname := strings.Replace(name, " ", "_", -1)
//formattedartist := strings.Replace(artist, " ", "_", -1)


 log.Println("dakhal")



 url := "http://ws.audioscrobbler.com/2.0/?method=chart.gettoptracks&api_key=33795d658c2750a8026eff986c4ed2e1&format=json"
 
    // fetch url
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln("Error fetching:", err)
    }
    // defer response close
    defer resp.Body.Close()

    // confirm we received an OK status
    if resp.StatusCode != http.StatusOK {
        log.Fatalln("Error Status not OK:", resp.StatusCode)
    }

    // read the entire body of the response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln("Error reading body:", err)
    }

    // create an empty instance of Feed struct
    // this is what gets filled in when unmarshaling JSON


    
    if err := json.Unmarshal(body, &entries); err != nil {
        log.Fatalln("Error decoing JSON", err)
    }

    
 log.Println("before for")
if len(entries.Tracks.Track) == 0 {
        log.Println("09876")
        return "", "", "", "","",fmt.Errorf("Artist or album not found")
    }else{

        res:= "Top 10 tracks are "
  

    for i := 0; i < 10; i++{
        entry := entries.Tracks.Track[i].Name;
        // log.Println(entries.TopAlbums.Album[i].Name)
        // log.Println(entries.TopAlbums.Album[i].Playcount)
        // log.Println(entry.Name)

        res = res+ entry + ", "




        //alb := album{Name: entries.TopAlbums.Album[i].Name, Playcount: entries.TopAlbums.Album[i].Playcount, Artist: entry.Name }
        //albums = append(albums, alb)


       

        

    }
     //log.Println("after for")

return "","" ,"", res, "", nil


}




}


func fetchtag(tag string) (string,string, string, string, string, error) {

var entries Feed3
 //formattedname := strings.Replace(name, " ", "_", -1)
//formattedartist := strings.Replace(artist, " ", "_", -1)






 url := "http://ws.audioscrobbler.com/2.0/?method=tag.gettoptracks&tag="+tag+"&api_key=33795d658c2750a8026eff986c4ed2e1&format=json"
 
    // fetch url
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln("Error fetching:", err)
    }
    // defer response close
    defer resp.Body.Close()

    // confirm we received an OK status
    if resp.StatusCode != http.StatusOK {
        log.Fatalln("Error Status not OK:", resp.StatusCode)
    }

    // read the entire body of the response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln("Error reading body:", err)
    }

    // create an empty instance of Feed struct
    // this is what gets filled in when unmarshaling JSON


    
    if err := json.Unmarshal(body, &entries); err != nil {
        log.Fatalln("Error decoing JSON", err)
    }

    
 log.Println("before for")
if len(entries.Tracks.Track) == 0 {
        log.Println("09876")
        return "", "", "", "", "",fmt.Errorf("Tag not found")
    }else{

        res:= "10 tracks related to this tag are "
  

    for i := 0; i < 10; i++{
        entry := entries.Tracks.Track[i].Name;
        // log.Println(entries.TopAlbums.Album[i].Name)
        // log.Println(entries.TopAlbums.Album[i].Playcount)
        // log.Println(entry.Name)

        res = res+ entry + ", "




        //alb := album{Name: entries.TopAlbums.Album[i].Name, Playcount: entries.TopAlbums.Album[i].Playcount, Artist: entry.Name }
        //albums = append(albums, alb)


       

        

    }
     log.Println("after for")

return "","" ,"", "", res, nil


}




}



func fetchfromalbum(name string, artist string) (string,string, string, string, string, error) {

var entries Feed2
 //formattedname := strings.Replace(name, " ", "_", -1)
//formattedartist := strings.Replace(artist, " ", "_", -1)


 formattedartist := artist
 formattedname := name

log.Println(formattedname)
log.Println(formattedartist)


 url := "http://ws.audioscrobbler.com/2.0/?method=album.getinfo&api_key=33795d658c2750a8026eff986c4ed2e1&artist="+formattedartist +"&album="+formattedname+"&format=json"
 
    // fetch url
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln("Error fetching:", err)
    }
    // defer response close
    defer resp.Body.Close()

    // confirm we received an OK status
    if resp.StatusCode != http.StatusOK {
        log.Fatalln("Error Status not OK:", resp.StatusCode)
    }

    // read the entire body of the response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln("Error reading body:", err)
    }

    // create an empty instance of Feed struct
    // this is what gets filled in when unmarshaling JSON


    
    if err := json.Unmarshal(body, &entries); err != nil {
        log.Fatalln("Error decoing JSON", err)
    }

    

if len(entries.Album.Tracks.Track) == 0 {
        log.Println("09876")
        return "", "", "","","",fmt.Errorf("Artist or album not found")
    }else{

        res:= "Track list is "
  

    for i := 0; i < len(entries.Album.Tracks.Track); i++{
        entry := entries.Album.Tracks.Track[i].Name;
        // log.Println(entries.TopAlbums.Album[i].Name)
        // log.Println(entries.TopAlbums.Album[i].Playcount)
        // log.Println(entry.Name)

        res = res+ entry + ", "




        //alb := album{Name: entries.TopAlbums.Album[i].Name, Playcount: entries.TopAlbums.Album[i].Playcount, Artist: entry.Name }
        //albums = append(albums, alb)


       

        

    }

return "","" ,res, "", "", nil


}




}


func fetch(name string) (string, string ,string, string, string, error) {


var entries Feed
 //formattedname := strings.Replace(name, " ", "_", -1)

formattedname := name

 url := "http://ws.audioscrobbler.com/2.0/?method=artist.gettopalbums&artist="+ formattedname + "&api_key=33795d658c2750a8026eff986c4ed2e1&format=json"
 albums:= [] album{}
    // fetch url
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln("Error fetching:", err)
    }
    // defer response close
    defer resp.Body.Close()

    // confirm we received an OK status
    if resp.StatusCode != http.StatusOK {
        log.Fatalln("Error Status not OK:", resp.StatusCode)
    }

    // read the entire body of the response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln("Error reading body:", err)
    }

    // create an empty instance of Feed struct
    // this is what gets filled in when unmarshaling JSON


    
    if err := json.Unmarshal(body, &entries); err != nil {
        log.Fatalln("Error decoing JSON", err)
    }

    

if len(entries.TopAlbums.Album) == 0 {
        return "", "", "","","",fmt.Errorf("Artist not found")
    }else{

        res:= "Top albums for this artist are "
  

    for i := 0; i < 5; i++{
        entry := entries.TopAlbums.Album[i].Artist;
        log.Println(entries.TopAlbums.Album[i].Name)
        log.Println(entries.TopAlbums.Album[i].Playcount)
        log.Println(entry.Name)

        res = res+ entries.TopAlbums.Album[i].Name + ", "




        alb := album{Name: entries.TopAlbums.Album[i].Name, Playcount: entries.TopAlbums.Album[i].Playcount, Artist: entry.Name }
        albums = append(albums, alb)


       

        

    }

return res,"" ,"", "", "", nil


}

}



func fetchtracks(name string) (string,string , string, string, string, error) {


var entries Feed1
 //formattedname := strings.Replace(name, " ", "_", -1)
formattedname := name

 url := "http://ws.audioscrobbler.com/2.0/?method=artist.gettoptracks&artist="+ formattedname + "&api_key=33795d658c2750a8026eff986c4ed2e1&format=json"
 tracks:= [] track{}
    // fetch url
    resp, err := http.Get(url)
    if err != nil {
        log.Fatalln("Error fetching:", err)
    }
    // defer response close
    defer resp.Body.Close()

    // confirm we received an OK status
    if resp.StatusCode != http.StatusOK {
        log.Fatalln("Error Status not OK:", resp.StatusCode)
    }

    // read the entire body of the response
    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatalln("Error reading body:", err)
    }

    // create an empty instance of Feed struct
    // this is what gets filled in when unmarshaling JSON


    
    if err := json.Unmarshal(body, &entries); err != nil {
        log.Fatalln("Error decoing JSON", err)
    }

    

if len(entries.TopTracks.Track) == 0 {
        return "", "","", "", "", fmt.Errorf("Artist not found")
    }else{
   res :="The top tracks are "

    for i := 0; i < 5; i++{

        res = res + entries.TopTracks.Track[i].Name +", "
        // entry := entries.TopTracks.Track[i].Artist;
        
        //log.Println(entries.TopTracks.Track[i].Playcount)
        //log.Println(entry.Name)

        trk := track{Name: entries.TopTracks.Track[i].Name, Playcount: entries.TopTracks.Track[i].Playcount }
        tracks = append(tracks, trk)


       

        

    }

return "", res,"", "", "", nil


}

}




func sampleProcessor(session Session, message string) (string,string ,string, string, string, error) {
   
    //log.Println(message[0:10])


if(len(message)>10){
    if(message[0:10] == "top albums"){
    log.Println("0")
    x,_,_,_,_,z:=fetch(message[11:len(message)])  

    return x, "", "","", "",z

     }

     if(message[0:10] == "top tracks"){
    log.Println("0")
    _,y,_,_,_,z:=fetchtracks(message[11:len(message)])  

    return "", y,"","", "", z

    }


}

if(strings.Contains(message,":")){
        
         res:= strings.Split(message,":")
         log.Println(res[0])
         log.Println(res[1])
         artist,album := res[0],res[1]

        _,_,w,_,_,z := fetchfromalbum(album,artist)
        return "","",w,"","",z
    }

if(message=="top ten"){
            //log.Println("my jigg")
        _,_,_,b,_,z := fetchchart()
        return "", "", "", b, "",z
            


        }

if(len(message)>3){        

if(message[0:3]=="tag"){
                _,_,_,_,c,z:=fetchtag(message[4:len(message)])  

    return "", "", "","", c,z
            }

        }


    return "", "","" , "", "",fmt.Errorf("Please enter a valid message")                 





   



    

    

}

// withLog Wraps HandlerFuncs to log requests to Stdout
func withLog(fn http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        c := httptest.NewRecorder()
        fn(c, r)
        log.Printf("[%d] %-4s %s\n", c.Code, r.Method, r.URL.Path)

        for k, v := range c.HeaderMap {
            w.Header()[k] = v
        }
        w.WriteHeader(c.Code)
        c.Body.WriteTo(w)
    }
}

// writeJSON Writes the JSON equivilant for data into ResponseWriter w
func writeJSON(w http.ResponseWriter, data JSON) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(data)
}


// ProcessFunc Sets the processor of the chatbot
func ProcessFunc(p Processor) {
     processor = p
 }

// handleWelcome Handles /welcome and responds with a welcome message and a generated UUID
func handleWelcome(w http.ResponseWriter, r *http.Request) {
    // Generate a UUID.
    hasher := md5.New()
    hasher.Write([]byte(strconv.FormatInt(time.Now().Unix(), 10)))
    uuid := hex.EncodeToString(hasher.Sum(nil))

    // Create a session for this UUID
    sessions[uuid] = Session{}

    // Write a JSON containg the welcome message and the generated UUID
    writeJSON(w, JSON{
        "uuid":    uuid,
        "message": WelcomeMessage,
    })
}

func handleChat(w http.ResponseWriter, r *http.Request) {
    // Make sure only POST requests are handled
    if r.Method != http.MethodPost {
        http.Error(w, "Only POST requests are allowed.", http.StatusMethodNotAllowed)
        return
    }

    // Make sure a UUID exists in the Authorization header
    uuid := r.Header.Get("Authorization")
    if uuid == "" {
        http.Error(w, "Missing or empty Authorization header.", http.StatusUnauthorized)
        return
    }

    // Make sure a session exists for the extracted UUID
    session, sessionFound := sessions[uuid]
    if !sessionFound {
        http.Error(w, fmt.Sprintf("No session found for: %v.", uuid), http.StatusUnauthorized)
        return
    }

    // Parse the JSON string in the body of the request
    data := JSON{}
    if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
        http.Error(w, fmt.Sprintf("Couldn't decode JSON: %v.", err), http.StatusBadRequest)
        return
    }
    defer r.Body.Close()

    // Make sure a message key is defined in the body of the request
    _, messageFound := data["message"]
    if !messageFound {
        http.Error(w, "Missing message key in body.", http.StatusBadRequest)
        return
    }


    // Process the received message
    message, tracks, tracklist , top, tag, err:= processor(session, data["message"].(string))
    log.Println(top+"snsnsns")
    log.Println(err)
    if err != nil {
        http.Error(w, err.Error(), 422 /* http.StatusUnprocessableEntity */)
        return
    }

    log.Println(top)

    if message != "" {
        writeJSON(w, JSON{
        "message": message,
    })
    

    }else{ 
        if tracks!="" {

       writeJSON(w, JSON{
        "message": tracks,
    }) 
    }else{

        if tracklist!=""{
            writeJSON(w, JSON{
        "message": tracklist,
    })

        }else{

            if top!=""{
                writeJSON(w, JSON{
        "message": top,
    })

            }else{
                //log.Println("0123")
           writeJSON(w, JSON{
        "message": tag,
    })
            }
             
        }

            
    }
    }
        
        
    
    

    

}

// handle Handles /
func handle(w http.ResponseWriter, r *http.Request) {
    body :=
        "<!DOCTYPE html><html><head><title>Chatbot</title></head><body><pre style=\"font-family: monospace;\">\n" +
            "Available Routes:\n\n" +
            "  GET  /welcome -> handleWelcome\n" +
            "  POST /chat    -> handleChat\n" +
            "  GET  /        -> handle        (current)\n" +
            "</pre></body></html>"
    w.Header().Add("Content-Type", "text/html")
    fmt.Fprintln(w, body)
}

// Engage Gives control to the chatbot
func Engage(addr string) error {
    // HandleFuncs
    mux := http.NewServeMux()
    mux.HandleFunc("/welcome", withLog(handleWelcome))
    mux.HandleFunc("/chat", withLog(handleChat))
    mux.HandleFunc("/", withLog(handle))
    //mux.HandleFunc("/fetch", withLog(fetch))

    // Start the server
    return http.ListenAndServe(addr, cors.CORS(mux))
}

func main() {
    // Uncomment the following lines to customize the chatbot
    WelcomeMessage = "Welcome to MaestroBot! Please enter a message in one of the following formats:<br> 1. To get top 5 albums for an artist: top albums (artist name) <br> 2. To get top tracks for an artist: top tracks (artist name)<br> 3. To get a specific artist's album's tracklist: (artist name):(album name)<br> 4. To get top 10 songs related to a specific tag: tag (tag name)"
    ProcessFunc(processor)

   

    // Use the PORT environment variable
    port := os.Getenv("PORT")
    // Default to 3000 if no PORT environment variable was defined
    if port == "" {
        port = "3000"
    }

    // Start the server
    fmt.Printf("Listening on port %s...\n", port)
    log.Fatalln(Engage(":" + port))
}

