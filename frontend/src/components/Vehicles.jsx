import { useEffect, useState } from "react";
import axios from "axios";

function Vehicles() {
  // State for storing the vehicle list.
  const [vehicles, setVehicles] = useState([]);

  // State for the vehicle form fields.
    const [customerId, setCustomerId] = useState("");
    const [regNo, setRegNo] = useState("");
    const [vehicleType, setVehicleType] = useState("");

    // State for storing customers used in the dropdown.
    const [customers, setCustomers] = useState([]);


    // Fetch all vehicles from the backend.
    const getVehicles = () => {
        axios
        .get("http://localhost:8080/api/vehicles")
        .then((response) => {
            setVehicles(response.data);
        });
    };

     // Fetch customers so they can be selected when adding a vehicle.
     const getCustomers = () => {
        axios
        .get("http://localhost:8080/api/customers")
        .then((response) => {
            setCustomers(response.data);
        });
    };

    // Fetch vehicles and customers when the component first loads.
    useEffect(() => {
        getVehicles();
        getCustomers();
    }, []);

    // Send the new vehicle to the backend.
    const handleSubmit = (e) => {
        e.preventDefault();
    
        if (!customerId || !regNo || !vehicleType) {
        alert("All fields are required");
        return;
        }
    
        axios
        .post("http://localhost:8080/api/vehicles", {
            customer_id: Number(customerId),
            reg_no: regNo,
            vehicle_type: vehicleType,
        })
        .then(() => {
            getVehicles();
            setCustomerId("");
            setRegNo("");
            setVehicleType("");
        });
    };

    return (
        <div>

            {/* Vehicle form: used to add a new vehicle. */}
            <form onSubmit={handleSubmit}>
            {/* Customer dropdown: lets the user select the vehicle owner. */}
            <select
                value={customerId}
                onChange={(e) => setCustomerId(e.target.value)}
            >
                <option value="">Select Customer</option>

                {customers.map((customer) => (
                <option key={customer.ID} value={customer.ID}>
                    {customer.Name}
                </option>
                ))}
            </select>

            {/* Registration number input. */}
            <input
                type="text"
                placeholder="Enter Registration Number"
                value={regNo}
                onChange={(e) => setRegNo(e.target.value)}
            />

            {/* Vehicle type input. */}
            <input
                type="text"
                placeholder="Enter Vehicle Type"
                value={vehicleType}
                onChange={(e) => setVehicleType(e.target.value)}
            />

            {/* Submit button for adding the vehicle. */}
            <button type="submit">Add Vehicle</button>
            </form>

        <h2>Vehicles</h2>

        {/* Display each vehicle from the vehicles array. */}
        {vehicles.map((vehicle) => (
            <div key={vehicle.ID}>
            <p>Registration No: {vehicle.RegNo}</p>
            <p>Vehicle Type: {vehicle.VehicleType}</p>
            {/*Display the customer name returned by the backend JOIN.*/}
                <p>Customer: {vehicle.CustomerName}</p>
            </div>
        ))}
        </div>
    );
 }

    export default Vehicles;