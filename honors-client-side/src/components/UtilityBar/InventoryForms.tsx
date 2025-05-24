import React, {useEffect, useState} from "react";



export function CreateForm(props: {id:string, onClose: () => void, UpdateRender: () => void}){
    return (
        <div id={props.id}>
            <form id={"CreateInput"} onSubmit={(event) => HandleSubmit(event, "AddEquipment", "POST", props.onClose, props.UpdateRender)}>
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

export function EditForm(props: {id:string, onClose: () => void, UpdateRender: () => void}){
    return (
        <div id={props.id}>
            <form id={"CreateInput"} onSubmit={(event) => HandleSubmit(event, "EditEquipment", "PUT", props.onClose, props.UpdateRender)}>

                <label> Equipment_ID
                    <input type={"text"} name={"id"} required/>
                </label>

                <label> New Model
                    <input type={"text"} name={"model"} required/>
                </label>

                <label> New Equipment Type
                    <Equipment_Type/>
                </label>

                <label> New Equipment Status
                   <Equipment_Status/>
                </label>

            </form>
        </div>
    )
}

export function TransferForm(props: { id: string, onClose: () => void , UpdateRender: () => void}) {
    return (
        <div id={props.id}>
            <form id={"CreateInput"} onSubmit={(event) => HandleSubmit(event, "TransferEquipment", "PUT", props.onClose, props.UpdateRender)}>
                <label> Equipment ID
                    <input type={"text"} name={"id"} required/>
                </label>

                <label> Room Name
                    <Room_Name/>
                </label>

            </form>
        </div>
    )
}

export function DeleteForm(props: {id:string, onClose: () => void, UpdateRender: () => void}){
    return (
        <div id={props.id}>
            <form id={"CreateInput"} onSubmit={(event) => HandleSubmit(event, "RemoveEquipment", "DELETE", props.onClose, props.UpdateRender)} >
                <label> Equipment ID
                    <input type={"text"} name={"id"} required/>
                </label>

            </form>
        </div>
    )

}

export function AuditForm(props: {id:string, onClose: () => void, UpdateRender: () => void}){
    type Message = {
     message: string;
    };
    const [Message, SetMessages] = useState<Message[]>([])


    useEffect(() => {
       fetch(`http://localhost:8080/API/GetAuditLogs`)
           .then(res => {
                if (!res.ok) throw new Error("Failed to Get Audit Logs");
                return res.json();
        })
           .then(data => SetMessages(data))
    }, [])


    return (
        <div id={props.id}>
            <form id={"CreateInput"}
                  onSubmit={(event) => HandleSubmit(event, "RemoveEquipment", "DELETE", props.onClose, props.UpdateRender)}>
                  <ul id={"Audit-Entries"}>
                      {Message.map(message => (
                 <li className={"Audit-Entry"}>
                    <p> {message.message} </p>
                 </li>
                      ))}
                  </ul>

            </form>
        </div>
    )
}


function HandleSubmit(event: React.FormEvent<HTMLFormElement>, endpoint: string, Method: string, onClose: () => void, UpdateRender: () => void) {
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
        onClose()
        UpdateRender()
      if (!res.ok) throw new Error("Failed to submit the form");
      return res.json();
    })
    .then((responseData) => {
      console.log("Success :) :", responseData);
    })
    .catch((error) => {
      console.error("Error submitting the form:", error);
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

function Room_Name() {
    return (
        <select name={"room_name"}>
            <option value="Honors 401A"> Honors 401A</option>

            <option value="Honors Kitchen"> Honors Kitchen</option>

            <option value="Office 1"> Office 1</option>

            <option value="Honors Warehouse"> Honors Warehouse</option>

            <option value="HON 305E"> HON 305E</option>

            <option value="HON 418E"> HON 418E</option>

            <option value="Office 2"> Office 2</option>

            <option value="HON 411B"> HON 411B</option>
        </select>
    )
}

function Building_Type() {

    return (
        <select name={"building_type"}>
            <option value="Warehouse"> Warehouse</option>

            <option value="Classroom"> Classroom</option>

            <option value="Office"> Office</option>

            <option value="Lab"> Lab</option>

            <option value="Other"> Other</option>
        </select>
    )

}

