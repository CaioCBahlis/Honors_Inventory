import React, {useState} from 'react';
import logo from './logo.svg';
import Rocky from './download-1.png'
import './App.css';

export default function App() {

  return (
    <div className="App">
      <UtilityBar/>
        <div id={"Main-Section"}>
            <HeaderInfo/>
            <MainDashboard/>
        </div>
    </div>
  );
}

function UtilityBar(){
  return (
      <div id={"UtilityBar"}>
         <LastEdit/>
          <EditButtons/>
      </div>
  )
}


function HeaderInfo(){
    return (
        <div id={"HeaderInfo"}>
            <h1 id={"Title"}> Judy Genshaft's Honors College Inventory </h1>
            <div id={"Stats"}>
                <div id={"TotalEquipment"}> </div>
                <div id={"ToFix"}></div>
            </div>
        </div>
    )
}

function MainDashboard(){
    return (
        <div id={"MainDashboard"}>



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
            <UtilityButtons text={"Create"}/>
            <UtilityButtons text={"Edit"}/>
            <UtilityButtons text={"Transfer"}/>
            <UtilityButtons text={"Remove"}/>
            <UtilityButtons text={"Logs"}/>
    </div>
    )
}

function UtilityButtons(props: {text: string, }){
    const [showModal, setModalView] = useState(false)

    return (
        <>
        <button className={"UtilityButton"} onClick={() => setModalView(true)} > {props.text} </button>
            {showModal && <CreateModal onClose={() => setModalView(false)}/>}
        </>
    )
}

function CreateModal({ onClose }: { onClose: () => void }){
    return (
        <div className={"Modal"} id={"Create"}>
            <div id={"Form"}>
                <h2 id={"Modal-Title"}> Add Equipment </h2>

                <div id={"CreateForm"}>

                </div>


                <div id={"Submit-Area"}>
                    <button className={"FormButton"} id={"Submit"}>
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










