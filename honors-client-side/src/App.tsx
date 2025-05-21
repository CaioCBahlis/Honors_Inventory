import React, {useState} from 'react';
import logo from './logo.svg';
import './App.css';
import UtilityBar from "./components/UtilityBar/Utility-bar";
import HeaderInfo from "./components/HeaderInfo/HeaderInfo";


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


function MainDashboard(){
    return (
        <div id={"MainDashboard"}>



        </div>
    )
}





