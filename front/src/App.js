
import React, { useEffect, useState } from 'react';
import { Container, Typography, List, ListItem, ListItemText, Paper } from '@mui/material';

import './App.css';
import axios from "axios";

const olympicColors = ['#0085C7', '#F4C300', '#009F3D', '#DF0024', '#EF7C00'];

function App() {
    const [messages, setMessages] = useState([]);
    const [polling, setPolling] = useState(false);

    const fetchMessages = async () => {
        try {
            const response = await axios.get('http://localhost:3000/leaderboard');
            setMessages(response.data);
        } catch (error) {
            console.error('Error fetching messages:', error.message);
        }
    };

    useEffect(() => {
        let interval;
        if (!polling) {
            setPolling(true)
            fetchMessages(); // initial fetch
            interval = setInterval(fetchMessages, 1000); // fetch every 5 seconds
            setPolling(false)
        }
        return () => clearInterval(interval);
    }, [polling]);

    return (
        <div className="App">
            <Container maxWidth="sm">
                <Typography variant="h4" gutterBottom>
                    Olympic Race Condition
                </Typography>
                <Paper elevation={3} className="Paper">
                    <List>
                        {Object.keys(messages).map((key, index) => (
                            <ListItem key={index}>
                                <ListItemText
                                    primary={`${key} - ${messages[key]}`}
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
