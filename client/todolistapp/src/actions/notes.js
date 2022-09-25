export const getNotes = async () =>{
    const response = await fetch('http://localhost:8080/note', {
        method: 'GET',
        'Content-Type': 'application/json'
    })
    const data = await response.json()
    console.log('notes',data)
    return data
}