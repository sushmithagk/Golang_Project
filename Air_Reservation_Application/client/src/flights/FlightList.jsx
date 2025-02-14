import { useEffect, useState } from "react";
import PageHeader from "../header/PageHeader";
import axios from "axios";

function FlightList() {
    // let [flights,setFlights]=useState([ //state ref ele, function element to set state
    //     // {id:'1010',number:'At 765',airline_name:'Air Inadia',source:'Mysore',destination:'Trichy',capacity:180,price:5000.0},
    //     // {id:'1020',number:'At 711',airline_name:'Air Inadia',source:'Trichy',destination:'Mysore',capacity:180,price:5000.0},
    //     // {id:'1030',number:'At 711',airline_name:'Air Inadia',source:'Trichy',destination:'Mysore',capacity:180,price:5000.0}
    // ]); 
    const [flights,setFlights]=useState([]);
    const readAllFlights= async()=>{
        try{
            const baseurl='http://localhost:8080'
            const response = await axios.get(`${baseurl}/flights`);
            setFlights(response.data);
        
        }catch{
            alert('Server error');
        }
    };
    useEffect(()=>{readAllFlights(); },[]);  //[]->after update    {}->before update(the data)
    const OnDelete = async (id) => {
        if(!confirm("Are you sure to delete?")){
            return;
        }
        try {
            const baseUrl = 'http://localhost:8080'
            const response = await axios.delete(`${baseUrl}/flights/${id}`);
            alert(response.data.message)
            readAllFlights();
        } catch(error) {
            alert('Server Error');
        }
    }
    return (
        <>
           
           {/* frontend design */}
            <PageHeader PageNumber={1}/>
            <h3>List of Flights</h3>
            <div className="container">
                <table className="table table-primary table-striped">
                    <thead className="table-dark">
                        <tr>
                            <th scope="col">Flight Number</th>
                            <th scope="col">Airline Name</th>
                            <th scope="col">Source</th>
                            <th scope="col">Destination</th>
                            <th scope="col"> Ticket Price</th>

                            <th></th>
                        </tr>
                    </thead>
                    <tbody>
                        { flights.map( (flight) => {
                            return(
                                <tr>
                                <th scope="row">{flight.id}</th>
                                <td>{flight.airline_name}</td>
                                <td>{flight.source}</td>
                                <td>{flight.destination}</td>
                                <td>{flight.price}</td>
                                <td>
                                    <a href={`/flights/edit/${flight.id}`} className="btn btn-warning">Edit Price</a>
                                    <button className="btn btn-danger"onClick={()=>{OnDelete(flight.id);}}>Delete</button>                             
                               </td>
                            </tr>
                            );
                        
                        } )
                        }

                     {/* <th scope="row">6E 151</th>
                        <td>Indigo</td>
                        <td>Hyderabad</td>
                        <td>Bengaluru</td>
                        <td>
                            <a href="/flights/edit/202111222" className="btn btn-warning">Edit Price</a>
                            <button className="btn btn-danger">Delete</button>
                        </td>
                        </tr> */}
                          

                        
                           
                    </tbody>
                </table>
            </div>
        </>
    );
}

export default FlightList;
