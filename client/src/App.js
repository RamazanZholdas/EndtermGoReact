import React from "react";
import "./App.css";
import {Container} from "semantic-ui-react";
import ToDoList from "./To-Do-List.js";
import Pomodoro from "./Pomodoro.js";

function App() {
    return (
        <div>
            <Container>
		        <Pomodoro />
                <ToDoList />
            </Container>
        </div>
    );
}

export default App

