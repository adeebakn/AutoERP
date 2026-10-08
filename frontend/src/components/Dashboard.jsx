import { useEffect,useState } from "react";
import axios from "axios";
import Card from "./Card";

function Dashboard()
{
    const [dashboard,setDashboard] = useState(null);

    useEffect(()=>{
        axios
        .get("http://localhost:8080/api/dashboard")
        .then((response)=>{
            setDashboard(response.data);
        });
    },[])


    if (!dashboard) {
        return <p>Loading...</p>;
      }

    return(
     <div>
        <h2>Dashboard</h2>

        <div className="dashboard-cards">
            <Card
                title="Pending Work orders"
                value={dashboard.pending_work_orders}
            />

            <Card
                title="In Progress"
                value={dashboard.in_progress_work_orders}
            />

            <Card
               title="Completed"
               value={dashboard.completed_work_orders}
            />

            <Card
                title="Unpaid Invoices"
                value={dashboard.unpaid_invoices}
            />

            <Card
                title="Total Revenue"
                value={`₹${dashboard.total_revenue}`}
            />

        </div>
     </div>
    );
}

export default Dashboard;