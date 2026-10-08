import "./Sidebar.css";
import {Link} from "react-router-dom";

function Sidebar() {
  return (
    <aside className="sidebar">
      <Link to="/">Dashboard</Link>
      <Link to="/customers">Customers</Link>
      <Link to="/vehicles">Vehicles</Link>
      <Link to="/services">Services</Link>
      <Link to="/quotations">Quotations</Link>
      <Link to="/work-orders">Work Orders</Link>
      <Link to="/invoices">Invoices</Link>
      <Link to="/payments">Payments</Link>
    </aside>
  );
}

export default Sidebar;