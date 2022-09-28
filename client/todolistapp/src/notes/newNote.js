import React, {useState} from "react";

const NewNote = () =>{
    const getDate = () => {
        let d = new Date();
        let y = d.getFullYear();
        let m = d.getMonth() + 1;
        let dd = d.getDate();
        let h = d.getHours();
        let mm = d.getMinutes();
        return y + '-' + m + '-' + dd + 'T' + h + ':' + mm + ':00';
    }

    const onChange = e =>{
        if (e.target.name === 'due'){
            let dueDate = e.target.value;
            dueDate += ':00'
            setFormData({...formData, [e.target.name]: dueDate});
            return;
        }
        setFormData({...formData, [e.target.name]: e.target.value})
    }

    const [formData, setFormData] = useState({
        title:'',
        body:'',
        createAt: getDate(),
        due: getDate()
    })
    return(
        <div style={{margin:'1.2rem'}}>
            <form onSubmit={e =>{
                e.preventDefault();
                //TODO: send post request
                console.log(formData)
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
                <input style={{display:'block', padding:'0.4rem', font:'1.2rem', border:'1px solid #ccc'}} type='datetime-local' id='due' name='due' value={formData.due} onChange={event => {onChange(event)}} min={formData.createAt} required/>
                <input style={{display:'block', padding:'0.4rem', font:'1.2rem', border:'1px solid #ccc'}} type='submit'/>
            </form>
        </div>
    )
}

export default NewNote;