import React, {useState} from "react";
import '../actions/notes'

const getTimeDiffrence = (due) => {
    let currTimeStamp = new Date().getTime();
    let dueTimeStamp = new Date(due).getTime();
    let tDiff = Math.floor(dueTimeStamp - currTimeStamp);
    let diff = '';
    let days = Math.floor((tDiff / 86_400_000));
    diff += days + ' days';
    tDiff -= days * 86_400_000
    let hours = Math.floor((tDiff / 3_600_000))
    diff += hours + ' hours';
    tDiff -= hours * 3_600_000
    let minutes = Math.floor((tDiff / 60_000))
    diff += minutes + 'minutes';

    return diff;
}


const NoteItem = (props) => {
    const [summary, setSummary] = useState(true);
    return (
        <div style={{maxHeight: '400px', maxWidth: '550px'}}>
            {
                summary ?
                    <div key={props.note.Id} style={{margin: '1px'}}>
                        Title: {props.note.Title} Time to finish:{getTimeDiffrence(props.note.Due)}
                            <button onClick={() => setSummary(!summary)}>Full Note</button>

                    </div>
                    :
                    <div key={props.note.Id} style={{margin: '1px', grid: 'flex'}}>
                        <p>Title: {props.note.Title}</p>
                        <p>Body: {props.note.Body}</p>
                        <p>Created at: {props.note.Create}</p>
                        <p>Due to: {props.note.Due}</p>
                        <p>Time to finish:{getTimeDiffrence(props.note.Due)}</p>
                        <button onClick={() => props.deleteing(props.note.Id)}>Finish</button>
                        <button onClick={() => setSummary(!summary)}>Exit</button>
                    </div>

            }
        </div>

    )
}

export default NoteItem;