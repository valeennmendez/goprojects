console.log("Conectado...")

function GetAllAppointments(){

    fetch(`http://localhost:8080/appointments`,{
        method: "GET",
        credentials: "include"
    })
        .then(response =>{
            if(!response.ok){
                console.error("ERROR")
            }
            return response.json()
        })
        .then(data =>{
            const tableBody = document.querySelector(".tabla tbody")
            tableBody.innerHTML = "";

            console.log(data)

            data.forEach(appointment =>{
                const row = document.createElement("tr")
                
                const fecha = new Date(appointment.Fecha)
                const fechaFormateada = fecha.toLocaleDateString()

                row.innerHTML = `
                <td>${appointment.Paciente.Email}</td>
                <td>${appointment.Paciente.Dni}</td>
                <td>${fechaFormateada}</td>
                <td>${appointment.Hora}</td>
                `
                tableBody.appendChild(row)
            })
        })
        .catch(error => console.error("error: ",error))
}

document.addEventListener("DOMContentLoaded", function(e){

    GetAllAppointments()

})