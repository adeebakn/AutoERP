import { BrowserRouter, Routes, Route } from "react-router-dom";
import Header from "./components/Header";
import Sidebar from "./components/Sidebar";
import Dashboard from "./components/Dashboard";
import Customers from "./components/Customers";
import Vehicles from "./components/Vehicles";
import Services from "./components/Services";
import Quotations from "./components/Quotations";
import WorkOrders from "./components/Work_Orders";
import Invoices from "./components/Invoices";
import Payments from "./components/Payments";


function App() {
  return (
    <BrowserRouter>
      <div className="app">
        <Header />

        <div className="layout">
          <Sidebar />

          <main className="content">

            <Routes>
              <Route path="/" element={<Dashboard />} />
              <Route path="/customers" element={<Customers />} />
              <Route path="/vehicles" element={<Vehicles />} />
              <Route path="/services" element={<Services />} />
              <Route path="/quotations" element={<Quotations />} />
              <Route path="/work-orders" element={<WorkOrders />} />
              <Route path="/invoices" element={<Invoices />} />
              <Route path="/payments" element={<Payments />} />
            </Routes>


          </main>
        </div>
      </div>
    </BrowserRouter>
  );
}

export default App;