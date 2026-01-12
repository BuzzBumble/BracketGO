import styles from './css/home.module.css'

export default function Home() {
  return (
    <div className="flex items-center justify-center bg-zinc-50 font-sans dark:bg-black w-full">
      <main className="flex min-h-screen w-full max-w-5xl flex-col items-center py-32 px-16 bg-gray-600 sm:items-start">
        <h3 className="text-5xl mb-32 text-blue-200">Brackets</h3>
        <h4 className="text-2xl mb-8 text-blue-200">Create a new Bracket</h4>
        <input className="placeholder:text-blue-200 border-solid border-gray-400 border-2 rounded-md mb-8" type="text" placeholder="Enter a name..."/>
        <div className={`flex flex-col items-center gap-6 text-center sm:items-start sm:text-left ${styles.cards}`}>
          
        </div>
      </main>
    </div>
  );
}
