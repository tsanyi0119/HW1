import React from "react";
import ProductContent from "./components/ProductContent";
import HomeContent from "./components/HomeContent";
import OrderContent from "./components/OrderContent";
import 'bootstrap/dist/css/bootstrap.min.css';
import BusinessNavbar from "./components/BusinessNavbar";
import { BrowserRouter as Router, Routes, Route } from 'react-router-dom';

function App() {
  return (
    <>
    <BusinessNavbar />
      <Router>
        <Routes>
            <Route exact path="/home" element={<HomeContent />} />
            <Route exact path="/product" element={<ProductContent />} />
            <Route exact path="/order" element={<OrderContent />} />
        </Routes>
      </Router>
    </>
  );
}

export default App;
