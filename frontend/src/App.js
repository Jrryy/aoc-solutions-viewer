// noinspection BadExpressionStatementJS

import './App.css';
import {useEffect, useState} from "react";
import {Routes, Route} from "react-router";
import CodePage from "./CodePage";
import {Link} from "react-router-dom";

function App() {
    const [links, setLinks] = useState([])

    useEffect(() => {
        fetch('http://localhost:8000/AoC-2021')
            .then(resp => resp.json())
            .then(elements => setLinks(elements))
    }, [])

    return (
        <div className="App flex justify-start flex-row text-xl text-white">
            <div className="flex-initial p-10">
                <ul>
                    {links.map(link =>
                        <li>
                            <Link to={link}>Day {link.slice(4)}</Link>
                        </li>
                    )}
                </ul>
            </div>
            <div className="p-10">
                <Routes>
                    {links.map(link =>
                        <Route path={link} element={<CodePage/>}/>
                    )}
                </Routes>
            </div>
        </div>
    );
}

export default App;
