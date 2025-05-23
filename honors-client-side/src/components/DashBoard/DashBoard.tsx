import React, {useEffect} from "react";
import {useState} from "react";

type Equipment = {
  id: string;
  model: string;
  equipment_type: string;
  equipment_status: string;
  equipment_room: string;
  equipment_room_type: string;
  inserted_at: string
};

export default function MainDashboard(props: {State: number}){
    return (
        <div id={"MainDashboard"}>
            <ItemList State={props.State}/>

        </div>
    )
}

function ItemList(props: {State: number}){
    const [Equipment, setEquipment] = useState<Equipment[]>([]);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
       fetch("http://localhost:8080/API/GetEquipments")
           .then(res => {
                if (!res.ok) throw new Error("Failed to fetch items");
                return res.json();
        })
           .then(data => setEquipment(data))
    }, [props.State])

    return (
        <>
            <div id={"EquipmentTitle"}>
                <p> Id</p>
                <p> Model </p>
                <p> Equipment Type</p>
                <p> Status </p>
                <p> Room Name </p>
                <p> Building Type</p>
            </div>

            <ul className={"EquipmentList"}>
            {Equipment.map(item => (
                <li key={item.id}>
                    <p> {item.id} </p>
                    <p> {item.model} </p>
                    <p> {item.equipment_type} </p>
                    <p> {item.equipment_status} </p>
                    <p> {item.equipment_room} </p>
                    <p> {item.equipment_room_type} </p>
                </li>
            ))}
        </ul>
        </>
    )
}
