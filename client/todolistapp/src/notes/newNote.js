import React, {useState} from "react";

const NewNote = ({setAddNote}) =>{
    const getDate = () => {
        let d = new Date();
        let y = d.getFullYear();
        let m = ('00' + (d.getMonth() + 1)).slice(-2);
        let dd = ('00' + d.getDate()).slice(-2);
        let h = ('00' + d.getHours()).slice(-2);
        let mm = ('00' + d.getMinutes()).slice(-2);
        return y + '/' + m + '/' + dd + ' ' + h + ':' + mm ;
    }

    const onChange = e =>{
        setFormData({...formData, [e.target.name]: e.target.value})
    }

    const [formData, setFormData] = useState({
        title:'',
        body:'',
        createAt: getDate(),
        due: getDate().replaceAll('/', '-').replace(' ', 'T')
    })
    return(
        <div style={{margin:'1.2rem'}}>
            <form onSubmit={async e =>{
                e.preventDefault();
                console.log(formData.createAt)
                formData.due = formData.due.replaceAll('-', '/').replace('T', ' ');
                setAddNote(false);
                const response = await fetch('http://localhost:8080/note', {
                    method: 'POST', // *GET, POST, PUT, DELETE, etc.
                    mode: 'no-cors', // no-cors, *cors, same-origin
                    headers: {
                        'Content-Type': 'application/json'
                    },
                    body: JSON.stringify(formData) // body data type must match "Content-Type" header
                })
                console.log(response)
            }}>
                <input style={{display:'block', padding:'0.4rem', font:'1.2rem', border:'1px solid #ccc'}} type='text' placeholder='Title' name='title' value={formData.title} onChange={event => {onChange(event)}} required/>
                <textarea
                    style={{display:'block', padding:'0.4rem', font:'1.2rem', border:'1px solid #ccc'}}
                    name='body'
                    cols='30'
                    rows='5'
                    placeholder='The body of your note'
                    value={formData.body}
                    onChange={event => {onChange(event)}}
                    required
                />
                <label htmlFor='due'>Due date:</label>
                <input style={{display:'block', padding:'0.4rem', font:'1.2rem', border:'1px solid #ccc'}} type='datetime-local' min='30-09-2022T00:00' id='due' name='due' value={formData.due} onChange={event => {onChange(event)}} required/>
                <input style={{display:'block', padding:'0.4rem', font:'1.2rem', border:'1px solid #ccc'}} type='submit'/>
            </form>
        </div>
    )
}

export default NewNote;