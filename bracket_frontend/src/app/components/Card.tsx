"use client"
import styles from '../css/card.module.css'

interface CardProps {
	name: string;
	id: number;
}

export default function Card(props: CardProps) {
	return (
		<li className={styles.card}>{props.name}</li>
	)
}