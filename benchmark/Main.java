import java.util.Random;
import java.time.Instant;

public class Main {

    public static void main(String[] args){
        int N = Integer.parseInt(args[0]);
        long startTime = Instant.now().toEpochMilli();
        int[][] A = new int[N][N];
	    int[][] B = new int[N][N];
	    int[][] C = new int[N][N];
        Random rand = new Random();
        // just like python init using 3 loop
        for (int i = 0; i < N; i++) {
            for (int j = 0; j < N; j++) {
                A[i][j] = rand.nextInt();
            }
        }
        for (int i = 0; i < N; i++) {
            for (int j = 0; j < N; j++) {
                B[i][j] = rand.nextInt();
            }
        }
        for (int i = 0; i < N; i++) {
            for (int j = 0; j < N; j++) {
                B[i][j] = rand.nextInt();
            }
        }
        for (int i = 0; i < N; i++) {
            for (int j = 0; j < N; j++) {
                for (int k = 0; k < N; k++) {
                    C[i][j] += A[i][k] * B[k][i];
                }
            }
        }

        long endTime = Instant.now().toEpochMilli();
 
        long timeElapsed = endTime - startTime;
 
        System.out.println("program took: " + timeElapsed + " ms");
    }
}