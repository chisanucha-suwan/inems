"use client";

import { useState } from "react";

export default function Test() {
    const [num, setNum] = useState(0);
    return (
        <button className="bg-red-800" onClick={() => { setNum(num + 1); console.info(num); }}>testttttttt</button>
    )
}