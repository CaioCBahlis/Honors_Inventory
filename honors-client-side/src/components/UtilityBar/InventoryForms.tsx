import React from "react";

export function CreateForm(props: {id:string}){
    return (
        <div id={props.id}>
            <form id={"CreateInput"} onSubmit={(event) => HandleSubmit(event, "AddEquipment", "POST")}>
                <label> Model
                    <input type={"text"} name={"model"} required/>
                </label>

                <label> Equipment Type
                        <Equipment_Type/>
                    </label>
            </form>
        </div>
    )
}

export function EditForm(props: {id:string}){
    return (
        <div id={props.id}>
            <form id={"CreateInput"} onSubmit={(event) => HandleSubmit(event, "EditEquipment", "PUT")}>

                <label> Equipment_ID
                    <input type={"text"} name={"id"} required/>
                </label>

                <label> Model
                    <input type={"text"} name={"model"} required/>
                </label>

                <label> Equipment Type
                    <Equipment_Type/>
                </label>

                <label> Equipment Status
                   <Equipment_Status/>
                </label>

            </form>
        </div>
    )
}

export function TransferForm(props: { id: string }) {
    return (
        <div id={props.id}>
            <form id={"CreateInput"} onSubmit={(event) => HandleSubmit(event, "TransferEquipment", "PUT")}>
                <label> Equipment ID
                    <input type={"text"} name={"id"} required/>
                </label>

                <label> Room Name
                    <input type={"text"} name={"room_name"} required/>
                </label>

                <label> Building Type
                    <Building_Type/>
                </label>

            </form>
        </div>
    )
}

export function DeleteForm(props: {id:string}){
    return (
        <div id={props.id}>
            <form id={"CreateInput"} onSubmit={(event) => HandleSubmit(event, "RemoveEquipment", "DELETE")} >
                <label> Equipment ID
                    <input type={"text"} name={"id"} required/>
                </label>

            </form>
        </div>
    )

}


function HandleSubmit(event: React.FormEvent<HTMLFormElement>, endpoint: string, Method: string) {
    event.preventDefault()
    const formData = new FormData(event.currentTarget)
    const data: Record<string, string> = {};

    formData.forEach((value, key) => {
        data[key] = value.toString();
    });

    const id = formData.get("id");
    let APIEndpoint: string = `http://localhost:8080/API/${endpoint}${id === null? "":`/${id}`}`


    fetch(APIEndpoint, {
        method: Method,
        headers: {"Content-Type": "application/json",},
        body: JSON.stringify(data)
    })
    .then((res) => {
      if (!res.ok) throw new Error("Failed to submit");
      return res.json();
    })
    .then((responseData) => {
      console.log("Success:", responseData);
    })
    .catch((error) => {
      console.error("Error:", error);
    });
}

function Equipment_Status(){
    return (
        <select name={"equipment_status"}>
            <option  value="Working"> Working</option>

            <option value="Needs Maintenance"> Needs Maintenance</option>

            <option value="Missing"> Missing</option>

            <option value="Broken"> Broken</option>
        </select>
    )
}

function Equipment_Type(){
     return (
         <select name={"equipment_type"}>
             <option value="Laptop"> Laptop</option>

             <option value="Desktop"> Desktop</option>

             <option value="Tablet"> Tablet </option>

             <option value="Monitor"> Monitor</option>

             <option value="Keyboard"> Keyboard</option>

             <option value="Mouse"> Mouse</option>

             <option value="Printer"> Printer</option>

             <option value="Projector"> Projector</option>

             <option value="Other"> Other</option>

         </select>
     )
}

function Building_Type(){

    return (
        <select name={"building_type"}>
            <option value="Warehouse"> Warehouse </option>

            <option value="Classroom"> Classroom </option>

            <option value="Office">  Office </option>

            <option value="Other">  Other </option>
        </select>
    )

}

