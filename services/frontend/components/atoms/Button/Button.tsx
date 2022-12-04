import styles from "./Button.module.css";

export type TButtonProps = {
  children: React.ReactNode;
};

const Button = ({ children }: TButtonProps) => {
  return (
    <button type="button" className={styles.button}>
      {children}
    </button>
  );
};

export default Button;
