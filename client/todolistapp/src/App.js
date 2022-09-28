import './App.css';
import axios from "axios";
import NoteItem from "./notes/noteItem";
import NewNote from "./notes/newNote";
import React, {useEffect, useState} from "react";

function App() {
    const [notes, setNotes] = useState(undefined)
    useEffect(() => {
        (async function () {
                try {
                    let data = await fetch('http://localhost:8080/note', {
                        method: 'GET',
                        'Content-Type': 'application/json'
                    })
                    let jsonresp = await data.json()
                    setNotes(jsonresp);
                } catch (e) {
                    console.error(e)
                }
            }
        )()
    }, [])

    async function deleteNote(id) {
        try {
            // let d = await fetch(`http://localhost:8080/note/${id}`, {
            //     method: 'DELETE',
            //     headers:{
            //         'Access-Control-Allow-Origin': '*',
            //         'Content-Type': 'application/json'
            //     }
            //
            // })
            // let js = await d.json();
            // console.log(js)
            await axios.delete(`http://localhost:8080/note/${id}`)
            let data = await fetch('http://localhost:8080/note', {
                method: 'GET',
                'Content-Type': 'application/json'
            })
            let jsonresp = await data.json()
            setNotes(jsonresp);
        } catch (e) {
            console.error(e)
        }
    }


    return (
        <div style={{display: "flex", backgroundColor: 'lightblue'}}>
            <div>
                {
                    notes ? notes.msg.map(note => {
                            return (
                                <NoteItem key={note.Id} note={note} deleteing={deleteNote}/>
                            )
                        })
                        : (<div>error</div>)

                }
            </div>
            <NewNote />
            <button style={{height: '20px', width: '75px'}}>Add Note</button>
        </div>
    );
}

export default App;
