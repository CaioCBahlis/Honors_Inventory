import React, {useState} from 'react';
import logo from './logo.svg';
import './App.css';
import UtilityBar from "./components/UtilityBar/Utility-bar";
import HeaderInfo from "./components/HeaderInfo/HeaderInfo";
import MainDashboard from "./components/DashBoard/DashBoard";


export default function App() {
    const [ItemState, ChangeItemState] = useState(0)
  return (
    <div className="App">
      <UtilityBar State={ItemState} UpdateState={() => ChangeItemState(ItemState => ItemState + 1)}/>
        <div id={"Main-Section"}>
            <HeaderInfo State={ItemState}/>
            <MainDashboard State={ItemState}/>
        </div>
    </div>
  );
}






