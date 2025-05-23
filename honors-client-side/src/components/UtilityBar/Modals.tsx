import React, {FC, SetStateAction, useState} from "react";

export default function CreateModal({ onClose, FormID, FormComponent, ModalTitle, UpdateRender}: { onClose: () => void, FormID: string, FormComponent: React.FC<{ id: string, onClose: () => void, UpdateRender: () => void }>, ModalTitle: string, UpdateRender: () =>void}){

    return (
        <div className={"Modal"} id={"Create"}>
            <div id={"Form"}>
                <h2 id={"Modal-Title"}> {ModalTitle} </h2>

                 <FormComponent id={FormID} onClose={onClose} UpdateRender={UpdateRender}/>

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





