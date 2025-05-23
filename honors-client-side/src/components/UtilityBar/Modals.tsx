import React, {FC, SetStateAction, useState} from "react";

export default function CreateModal({ onClose, FormID, FormComponent, ModalTitle, SetState, State}: { onClose: () => void, FormID: string, FormComponent: React.FC<{id: string}>, ModalTitle: string, SetState: React.Dispatch<SetStateAction<number>>, State:number}){

    return (
        <div className={"Modal"} id={"Create"} onSubmit={() => HandleSubmit(onClose, SetState, State)}>
            <div id={"Form"}>
                <h2 id={"Modal-Title"}> {ModalTitle} </h2>

                 <FormComponent id={FormID}/>

                <div id={"Submit-Area"}>
                    <button className={"FormButton"} id={"Submit"} form={"CreateInput"}>
                       Submit
                    </button>

                    <button className={"FormButton"} id={"Cancel"} onClick={onClose}>
                        Cancel
                    </button>
                </div>
            </div>
        </div>
    )
}

function HandleSubmit(SetSubmit: React.Dispatch<SetStateAction<boolean>>, SetState: React.Dispatch<SetStateAction<number>>, State: number){
    SetSubmit(false)
    SetState(State => State + 1)
}





