import React, {Component, SetStateAction, useEffect, useState} from "react";
import Rocky from "./download-1.png";
import CreateModal, {ROMModal} from "./Modals";
import {CreateForm, EditForm, TransferForm, DeleteForm, AuditForm} from "./InventoryForms";



export default function UtilityBar(props: {State: number, UpdateState: () => void}){
  return (
      <div id={"UtilityBar"}>
          <LastEdit State={props.State}/>
          <EditButtons UpdateState={props.UpdateState}/>
      </div>
  )
}

function LastEdit(props: {State: number}){
    const [Time, SetTime] = useState("")


     useEffect(() => {
        const fetchData = async () => {
            try {
                const res = await fetch("http://localhost:8080/API/GetLastInsert",
                    {cache: "no-cache"});
                const json = await res.json();
                SetTime(json);
            } catch (err) {
                console.error("Error fetching data", err);
            }
        };

        fetchData();
    }, [props.State]);



    return (
       <div id={"LastEdit"}>
           <h3> Last Edit </h3>
           <img id="LastEditImg" src={Rocky}/>
           <h3> Rocky The Bull! </h3>
           <h3> {Time}</h3>
           <hr/>
       </div>
    )
}

function EditButtons(props: {UpdateState: () => void}){
    return (
        <div id={"EditButtons"}>
            <UtilityButtons text={"Create"} Modal={CreateModal} MyFormComponent={CreateForm} MyFormId={"CreateForm"} ModalTitle={"Add Equipment"} UpdateRender={props.UpdateState}/>
            <UtilityButtons text={"Edit"} Modal={CreateModal} MyFormComponent={EditForm} MyFormId={"CreateForm"} ModalTitle={"Edit Equipment"} UpdateRender={props.UpdateState}/>
            <UtilityButtons text={"Transfer"} Modal={CreateModal} MyFormComponent={TransferForm} MyFormId={"CreateForm"} ModalTitle={"Transfer Equipment"} UpdateRender={props.UpdateState}/>
            <UtilityButtons text={"Delete"} Modal={CreateModal} MyFormComponent={DeleteForm} MyFormId={"CreateForm"} ModalTitle={"Remove Equipment"} UpdateRender={props.UpdateState}/>
            <UtilityButtons text={"Audit Logs"} Modal={ROMModal}  MyFormComponent={AuditForm} MyFormId={"ROMForm"} ModalTitle={"Audit Logs"} UpdateRender={props.UpdateState}/>
    </div>
    )
}

function UtilityButtons(props: {text: string, Modal: React.FC<any>, MyFormComponent: React.FC<{id:string, onClose: () => void, UpdateRender: () => void}>, MyFormId: string, ModalTitle: string, UpdateRender: () => void}){
    const [showModal, setModalView] = useState(false)

    return (
        <>
        <button className={"UtilityButton"} onClick={() => setModalView(true)} > {props.text} </button>
            {showModal && <props.Modal onClose={() => setModalView(false)} UpdateRender={props.UpdateRender} FormID={props.MyFormId} FormComponent={props.MyFormComponent} ModalTitle={props.ModalTitle}/>}
        </>
    )
}

