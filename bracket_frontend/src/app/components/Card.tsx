"use client"
import styles from '../css/card.module.css'

interface CardProps {
	name: string;
	key: number;
}

export default function Card(props: CardProps) {
	return (
		<li className={styles.card} key={props.key}>{props.name}</li>
	)
}