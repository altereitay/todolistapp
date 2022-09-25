import React, {useEffect, useState} from "react";
import '../actions/notes'
import {getNotes} from "../actions/notes";

const NoteItem = () => {
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
        <div>
            {
                notes ? notes.msg.map(note => {
                    console.log(typeof note.Create)
                    return (
                        <div key={note.Id}>
                            <h1>title: {note.Title}</h1>
                            <h2>body: {note.Body}</h2>
                            <h3>created at: {note.Create}</h3>
                            <h4>due to: {note.Due}</h4>
                            <h5>time to finish:</h5>
                            <button>finish</button>
                        </div>
                    )
                })
                    :(<div>error</div>)

            }

        </div>
    )
}

export default NoteItem