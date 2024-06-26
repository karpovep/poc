// src/App.js
import React, { useEffect, useState } from 'react';
import { Container, Typography, List, ListItem, ListItemText, Paper } from '@mui/material';
import { CloudClient } from './protos/main_grpc_web_pb.js';
import { CloudObject } from './protos/main_pb';

import './App.css';

const olympicColors = ['#0085C7', '#F4C300', '#009F3D', '#DF0024', '#EF7C00'];

function App() {
    const [messages, setMessages] = useState([]);

        const client = new CloudClient('http://localhost:8080');
        const stream = client.subscribe(new CloudObject(), {});

        stream.on('data', response => {
            setMessages(prevMessages => [...prevMessages, response.toObject()]);
        });

        stream.on('error', err => {
            console.error('Error:', err.message);
        });

    return (
        <div className="App">
            <Container maxWidth="sm">
                <Typography variant="h4" gutterBottom>
                    Olympic Race Condition
                </Typography>
                <Paper elevation={3} className="Paper">
                    <List>
                        {messages.map((message, index) => (
                            <ListItem key={index}>
                                <ListItemText
                                    primary={`${message.getName()} - ${message.getTimefinish()}`}
                                    style={{ color: olympicColors[index % olympicColors.length] }}
                                />
                            </ListItem>
                        ))}
                    </List>
                </Paper>
            </Container>
        </div>
    );
}

export default App;
