import React from "react"
import { useState } from "react";
import { useEffect } from "react";
import ReactPlayer from "react-player"

export default function RecCard() {

    const [error, setError] = useState(null);
    const [isLoaded, setIsLoaded] = useState(false);
    const [songs, setSongs] = useState([]);

    useEffect(()=>{
        fetch("http://localhost:3001/songs").
        then(res => res.json()).
        then((result) => {
            setIsLoaded(true);
            setSongs(result);
        },
        (error) => {
            setIsLoaded(true);
            setError(error);
        })
    }, []);

    if(error){
        return <h1>{error.message}</h1>
    } else if (!isLoaded){
        return <h1> Loading . . . </h1>
    } else {
        return <ul>{songs.map(song => (
            <li key = {song.id}>
               {song.title}:
               <ReactPlayer url={song.link} controls={true}/></li>
        ))}</ul>
    }
};