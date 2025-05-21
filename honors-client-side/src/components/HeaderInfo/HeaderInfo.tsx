import React from "react";
import {useEffect, useState} from "react";

export default function HeaderInfo(){
    const [data, setData] = useState<any>(0);
    const [loading, setLoading] = useState(true);

    useEffect(() => {
        const fetchData = async () => {
            try {
                const res = await fetch("http://localhost:8080/API/GetEquipmentInfo");
                const json = await res.json();
                setData(json);
            } catch (err) {
                console.error("Error fetching data", err);
            } finally {
                setLoading(false);
            }
        };

        fetchData();
    }, []);


    return (
        <div id={"HeaderInfo"}>
            <h1 id={"Title"}> Judy Genshaft's Honors College Inventory </h1>

            <div id={"Stats"}>
                <div id={"TotalEquipment"}>


                    <div className={"Stats-Info"} id={"Total"}>
                        <h2 className={"Stats-Data"}> {data[0]} </h2>
                        <p className={"Stats-Desc"}> Total Items in Inventory </p>

                    </div>


                    <div className={"Stats-Info"} id={"ToFix"}>
                        <h2 className={"Stats-Data"}> {data[2]} </h2>
                        <p className={"Stats-Desc"}> Items Broken & Maintenance </p>

                    </div>


                    <div className={"Stats-Info"} id={"InWarehouse"}>
                        <h2 className={"Stats-Data"}> {data[1]} </h2>
                        <p className={"Stats-Desc"}> Items Stored in the Warehouse </p>

                    </div>
                </div>
            </div>

        </div>
    )
}