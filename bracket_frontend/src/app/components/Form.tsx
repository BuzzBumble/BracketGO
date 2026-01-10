"use client"
import { useState, FormEvent } from 'react';
import axios from 'axios';

export default function Form() {
	const [name, setName] = useState('');
	const [loading, setLoading] = useState(false);
	const [error, setError] = useState<string | null>(null);
	const [success, setSuccess] = useState(false);
	const baseURL = process.env.NEXT_PUBLIC_API_URL;

	const handleSubmit = async (e: FormEvent) => {
		e.preventDefault();
		setLoading(true);
		setError(null);
		setSuccess(false);

		try {
			await axios.post(baseURL + 'api/go/brackets', {
				name
			});

			setSuccess(true);
			setName('');
		} catch (err: any) {
			setError(err?.response?.data?.message || 'Something went wrong');
		} finally {
			setLoading(false);
		}
	}

	return (
		<form onSubmit={handleSubmit}>
			<h2>Test Form</h2>
			<div>
				<label>Name</label>
				<input type="text" value={name} onChange={(e) => setName(e.target.value)} required/>

			</div>
			<button type="submit" disabled={loading}>
				{loading ? 'Submitting...' : 'Submit'}
			</button>

			{success && <p style={{color : 'green'}}>Submitted</p>}
			{error && <p style={{color: 'red'}}>{error}</p>}
		</form>
	)
}