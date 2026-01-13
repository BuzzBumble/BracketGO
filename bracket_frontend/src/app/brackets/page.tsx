"use client"

import {useState, useEffect} from 'react';
import axios from 'axios'
import Card from '../components/Card';
import styles from '../css/home.module.css'

interface Bracket {
	id: number;
	name: string;
}

export default function BracketsIndex () {
	const [brackets, setBrackets] = useState([]);

	const [loading, setLoading] = useState(true);

	const [error, setError] = useState("");

	const apiReqURL: string | undefined = process.env.NEXT_PUBLIC_API_URL

	useEffect(() => {
		const fetchBrackets = async() => {
			if (apiReqURL == undefined) {
				setError("API Request URL is undefined");
				return
			}
			try {
				const response = await axios.get(apiReqURL + "brackets")
				setBrackets(response.data.data);
				setError("");
			} catch (err:any) {
				setError(err.message);
				setBrackets([]);
			} finally {
				setLoading(false);
			}
		}

		fetchBrackets();
	}, [])

	function renderBrackets() {
		return brackets.map((b: Bracket) => (
			<Card name={b.name} key={b.id} id={b.id}/>
		))
	}

	return (
		<div>
			<h1 className="text-5xl">All Brackets</h1>
			{error != "" ? <h1>{error}</h1> :
			<ul className={`flex flex-col items-center gap-6 text-center sm:items-start sm:text-left ${styles.cards}`}>
				{renderBrackets()}
			</ul>}
		</div>
		
	);
}