import React from 'react';
import './App.css';
import {Spiral} from "./components/Spiral";

export function App() {
    return (
        <div className="container w-screen">
            <p className="text-5x1">Fibonnaci Spiral</p>
            <br/>
            <div>
                <Spiral/>
            </div>
        </div>
    );
}
