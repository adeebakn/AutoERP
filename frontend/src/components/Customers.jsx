import { useEffect, useState } from "react";
import axios from "axios";

function Customers() {
  // State for storing the customer list and the customer currently being edited.
  const [customers, setCustomers] = useState([]);
  const [editingCustomer, setEditingCustomer] = useState(null);

  // State for storing the values entered in the customer form.
  const [name, setName] = useState("");
  const [phone, setPhone] = useState("");
  const [email, setEmail] = useState("");

  // Fetch all customers from the backend.
  const getCustomers = () => {
    axios
      .get("http://localhost:8080/api/customers")
      .then((response) => {
        setCustomers(response.data);
      });
  };

  // Fetch customers when the component first loads.
  useEffect(() => {
    getCustomers();
  }, []);

  // Fill the form with the selected customer's data when editing starts.
  useEffect(() => {
    if (editingCustomer) {
      setName(editingCustomer.Name);
      setPhone(editingCustomer.Phone);
      setEmail(editingCustomer.Email);
    }
  }, [editingCustomer]);

  // Clear the form after asking the user for confirmation.
  const handleCancel = () => {
    const confirmCancel = window.confirm(
      "Are you sure you want to clear the form?"
    );

    if (confirmCancel) {
      setEditingCustomer(null);
      setName("");
      setPhone("");
      setEmail("");
    }
  };

  // Handle both creating a new customer and updating an existing customer.
  const handleSubmit = (e) => {
    e.preventDefault();

    if (!name || !phone) {
      alert("Name and phone are required");
      return;
    }

    // Update the existing customer when editingCustomer contains a customer.
    if (editingCustomer) {
      axios
        .put(
          `http://localhost:8080/api/customers/${editingCustomer.ID}`,
          {
            name: name,
            phone: phone,
            email: email,
          }
        )
        .then((response) => {
          console.log(response.data);
          getCustomers();
          setEditingCustomer(null);
          setName("");
          setPhone("");
          setEmail("");
        });
    } else {
      // Create a new customer when no customer is currently being edited.
      axios
        .post("http://localhost:8080/api/customers", {
          name: name,
          phone: phone,
          email: email,
        })
        .then((response) => {
          console.log(response.data);
          setCustomers((previousCustomers) => [
            ...previousCustomers,
            response.data,
          ]);
          setName("");
          setPhone("");
          setEmail("");
        });
    }
  };

  const handleDelete = (id) =>{
    const confirmDelete = window.confirm("Are you sure you want to delete this customer? ");

    if(confirmDelete){
        axios
        .delete(`http://localhost:8080/api/customers/${id}`)
        .then(()=> {
            getCustomers();
        });
    }
  };
  

  return (
    <div>
      <h2>Customers</h2>

      {/* Customer form: used for both adding and editing customers. */}
      <form onSubmit={handleSubmit}>
        {/* Name input: displays the name state and updates it when the user types. */}
        <input
          type="text"
          placeholder="Enter Name"
          value={name}
          onChange={(e) => setName(e.target.value)}
        />

        {/* Phone input: displays the phone state and updates it when the user types. */}
        <input
          type="text"
          placeholder="Enter Phone"
          value={phone}
          onChange={(e) => setPhone(e.target.value)}
        />

        {/* Email input: displays the email state and updates it when the user types. */}
        <input
          type="text"
          placeholder="Enter Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />

        {/* Submit button: changes between Add and Update depending on edit mode. */}
        <button type="submit">
          {editingCustomer ? "Update Customer" : "Add Customer"}
        </button>

        {/* Cancel button: clears the form without submitting it. */}
        <button type="button" onClick={handleCancel}>
          Cancel
        </button>
      </form>

      <h3>Customer List</h3>

      {/* Display each customer from the customers array. */}
      {customers.map((customer) => (
        <div key={customer.ID}>
          <p>Name: {customer.Name}</p>
          <p>Phone: {customer.Phone}</p>
          <p>Email: {customer.Email}</p>

          {/* Edit button: stores the selected customer in editingCustomer. */}
          <button onClick={() => setEditingCustomer(customer)}>
            Edit
          </button>

          <button onClick={() => handleDelete(customer.ID)}>
            Delete
          </button>

        </div>
      ))}
    </div>
  );
}

export default Customers;