import React, {Component, SetStateAction, useEffect, useState} from "react";
import Rocky from "./download-1.png";
import CreateModal from "./Modals";
import {CreateForm, EditForm, TransferForm, DeleteForm} from "./InventoryForms";

export default function UtilityBar(props: {State: number,SetState: React.Dispatch<SetStateAction<number>>}){
  return (
      <div id={"UtilityBar"}>
          <LastEdit State={props.State}/>
          <EditButtons SetState={props.SetState} State={props.State}/>
      </div>
  )
}

function LastEdit(props: {State: number}){
    const [Time, SetTime] = useState("")


     useEffect(() => {
        const fetchData = async () => {
            try {
                const res = await fetch("http://localhost:8080/API/GetLastInsert");
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

function EditButtons(props: {State:number, SetState: React.Dispatch<SetStateAction<number>>}){
    return (
        <div id={"EditButtons"}>
            <UtilityButtons text={"Create"} MyFormComponent={CreateForm} MyFormId={"CreateForm"} ModalTitle={"Add Equipment"} SetState={props.SetState} State={props.State}/>
            <UtilityButtons text={"Edit"} MyFormComponent={EditForm} MyFormId={"CreateForm"} ModalTitle={"Edit Equipment"} SetState={props.SetState} State={props.State}/>
            <UtilityButtons text={"Transfer"} MyFormComponent={TransferForm} MyFormId={"CreateForm"} ModalTitle={"Transfer Equipment"} SetState={props.SetState} State={props.State}/>
            <UtilityButtons text={"Delete"} MyFormComponent={DeleteForm} MyFormId={"CreateForm"} ModalTitle={"Remove Equipment"} SetState={props.SetState} State={props.State}/>
            <UtilityButtons text={"Logs"} MyFormComponent={CreateForm} MyFormId={"CreateForm"} ModalTitle={"Audit Logs"} SetState={props.SetState} State={props.State}/>
    </div>
    )
}

function UtilityButtons(props: {text: string, MyFormComponent: React.FC<{id: string}>, MyFormId: string, ModalTitle: string,  SetState: React.Dispatch<SetStateAction<number>>, State:number }){
    const [showModal, setModalView] = useState(false)

    return (
        <>
        <button className={"UtilityButton"} onClick={() => setModalView(true)} > {props.text} </button>
            {showModal && <CreateModal onClose={() => setModalView(false)} FormID={props.MyFormId} FormComponent={props.MyFormComponent} ModalTitle={props.ModalTitle} SetState={props.SetState} State={props.State} />}
        </>
    )
}

