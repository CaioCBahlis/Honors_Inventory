import React, {Component, useState} from "react";
import Rocky from "./download-1.png";
import CreateModal from "./Modals";
import {CreateForm, EditForm, TransferForm, DeleteForm} from "./InventoryForms";


export default function UtilityBar(){
  return (
      <div id={"UtilityBar"}>
          <LastEdit/>
          <EditButtons/>
      </div>
  )
}

function LastEdit(){
    return (
       <div id={"LastEdit"}>
           <h3> Last Edit </h3>
           <img id="LastEditImg" src={Rocky}/>
           <h3> Rocky The Bull! </h3>
           <h3> 10:59 May, 19</h3>
           <hr/>
       </div>
    )
}

function EditButtons(){
    return (
        <div id={"EditButtons"}>
            <UtilityButtons text={"Create"} MyFormComponent={CreateForm} MyFormId={"CreateForm"} ModalTitle={"Add Equipment"}/>
            <UtilityButtons text={"Edit"} MyFormComponent={EditForm} MyFormId={"CreateForm"} ModalTitle={"Edit Equipment"}/>
            <UtilityButtons text={"Transfer"} MyFormComponent={TransferForm} MyFormId={"CreateForm"} ModalTitle={"Transfer Equipment"}/>
            <UtilityButtons text={"Delete"} MyFormComponent={DeleteForm} MyFormId={"CreateForm"} ModalTitle={"Remove Equipment"}/>
            <UtilityButtons text={"Logs"} MyFormComponent={CreateForm} MyFormId={"CreateForm"} ModalTitle={"Audit Logs"}/>
    </div>
    )
}

function UtilityButtons(props: {text: string, MyFormComponent: React.FC<{id: string}>, MyFormId: string, ModalTitle: string }){
    const [showModal, setModalView] = useState(false)

    return (
        <>
        <button className={"UtilityButton"} onClick={() => setModalView(true)} > {props.text} </button>
            {showModal && <CreateModal onClose={() => setModalView(false)} FormID={props.MyFormId} FormComponent={props.MyFormComponent} ModalTitle={props.ModalTitle}  />}
        </>
    )
}

