package dev.itsyourap.cynox;

import android.os.Bundle;
import android.widget.Button;
import android.widget.EditText;
import android.widget.TextView;

import androidx.activity.EdgeToEdge;
import androidx.appcompat.app.AppCompatActivity;
import androidx.core.graphics.Insets;
import androidx.core.view.ViewCompat;
import androidx.core.view.WindowInsetsCompat;

public class MainActivity extends AppCompatActivity {

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        EdgeToEdge.enable(this);
        setContentView(R.layout.activity_main);
        ViewCompat.setOnApplyWindowInsetsListener(findViewById(R.id.main), (v, insets) -> {
            Insets systemBars = insets.getInsets(WindowInsetsCompat.Type.systemBars());
            v.setPadding(systemBars.left, systemBars.top, systemBars.right, systemBars.bottom);
            return insets;
        });

        EditText flagEditText = findViewById(R.id.flagEditText);
        Button submitButton = findViewById(R.id.submitButton);
        TextView messageTextView = findViewById(R.id.messageTextView);

        submitButton.setOnClickListener(v -> {
            String flag = flagEditText.getText().toString();
            if (flag.equals(getString(R.string.flag))) {
                messageTextView.setText(getResources().getText(R.string.correct_flag));
            } else {
                messageTextView.setText(getResources().getText(R.string.incorrect_flag));
            }
        });
    }
}