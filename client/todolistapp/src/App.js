import './App.css';
import NoteItem from "./notes/noteItem";
import React, {useEffect, useState} from "react";

function App() {
    const [notes, setNotes] = useState(undefined)
    useEffect(() => {
        (async function() {
                try {
                    let data = await fetch('http://localhost:8080/note', {
                        method: 'GET',
                        'Content-Type': 'application/json'
                    })
                    let jsonresp = await data.json()
                    setNotes(jsonresp);
                }catch (e) {
                    console.error(e)
                }
            }
        )()
    }, [])

  return (
      <div style={{display:"flex", backgroundColor:'lightblue'}}>
          <div>
          {
              notes ? notes.msg.map(note => {
                  console.log(note)
                      return (
                          <NoteItem note = {note} />
                      )
                  })
                  :(<div>error</div>)

          }
          </div>
          <button style={{height: '20px', width: '75px'}}>Add Note</button>
      </div>
  );
}

export default App;
