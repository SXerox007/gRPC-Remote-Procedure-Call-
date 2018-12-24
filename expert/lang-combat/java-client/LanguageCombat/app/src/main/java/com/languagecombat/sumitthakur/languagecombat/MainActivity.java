package com.languagecombat.sumitthakur.languagecombat;

import android.annotation.SuppressLint;
import android.os.AsyncTask;
import android.support.v7.app.AppCompatActivity;
import android.os.Bundle;
import android.util.Log;

import java.util.concurrent.TimeUnit;

import io.grpc.ManagedChannel;
import io.grpc.ManagedChannelBuilder;

import io.grpc.okhttp.OkHttpChannelBuilder;
import login.Login;
import login.LoginServiceGrpc;


public class MainActivity extends AppCompatActivity {

//    private  ManagedChannel channel;
//    private  LoginServiceGrpc.LoginServiceBlockingStub loginServiceStub;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);
        //new GrpcTask().execute();
        languageCombat();
    }

    private void languageCombat() {
        //in host your public IP
        ManagedChannel channel = OkHttpChannelBuilder.forAddress("192.168.0.106", 50051)
                .usePlaintext()
                .build();
        connectionCall(channel);
    }

    private void connectionCall(ManagedChannel channel) {
        LoginServiceGrpc.LoginServiceBlockingStub loginClient = LoginServiceGrpc.newBlockingStub(channel);

        Login.LoginRequestData data = Login.LoginRequestData.newBuilder()
                .setEmailOrPhone("sumit@gmail.com")
                .setPassword("12345678")
                .build();

        Login.LoginRequest req = Login.LoginRequest.newBuilder().
                setLoginRequestData(data)
                .build();
        Log.e("Data Request",req.toString());
        Login.LoginResponse resp = loginClient.login(req);

        Log.e("Whole data:",resp.toString());
    }


//   MainActivity(){}
//
//    public void shutdown() throws InterruptedException {
//        channel.shutdown().awaitTermination(5, TimeUnit.SECONDS);
//    }

//    private void init(){
//        Log.d("Start","Hell0");
//        MainActivity client = new MainActivity("localhost", 50051);
//        try {
//            Log.d("Start",client.toString());
//            client.loginResquest();
//        }finally {
//            try {
//                client.shutdown();
//            } catch (InterruptedException e) {
//                e.printStackTrace();
//            }
//        }
//    }

//
//    private void loginResquest(){
//        Log.d("Login Request","Login request");
//        Login.LoginRequestData data = Login.LoginRequestData.newBuilder().setEmailOrPhone("blackdream@gmail.com").setPassword("12345678").build();
//        Login.LoginRequest req = Login.LoginRequest.newBuilder().setLoginRequestData(data).build();
//        Login.LoginResponse resp;
//        try{
//           // resp.getLoginResponseData();
//            resp = loginServiceStub.login(req);
//        } catch (StatusRuntimeException e) {
//
//        }
//        channel.shutdown();
//    }




//     MainActivity(ManagedChannel channel) {
//        this.channel = channel;
//        this.loginServiceStub = LoginServiceGrpc.newBlockingStub(channel);
//
//    }
//
//
//     MainActivity(String localhost, int port) {
//        this(ManagedChannelBuilder.forAddress(localhost, port)
//                .usePlaintext()
//                .build());
//
//    }
//
//


    @SuppressLint("StaticFieldLeak")
    private class GrpcTask extends AsyncTask<Void, Void, String> {

        private ManagedChannel mChannel;

        @Override
        protected void onPreExecute() {

        }

        @Override
        protected String doInBackground(Void... voids) {
            try {
                Log.e("Step 1","start");
                mChannel = OkHttpChannelBuilder
                        .forAddress("localhost", 5051)
                        .usePlaintext()
                        .build();
                Log.e("Is Shutdown", String.valueOf(mChannel.isShutdown()));
                Log.e("Step 2",mChannel.toString());

                LoginServiceGrpc.LoginServiceBlockingStub stub =
                        LoginServiceGrpc.newBlockingStub(mChannel);
                Log.e("Step 3",stub.toString());
                Login.LoginRequestData data = Login.LoginRequestData.newBuilder().setEmailOrPhone("sumit@gmail.com").setPassword("12345678").build();
                Login.LoginRequest req = Login.LoginRequest.newBuilder().setLoginRequestData(data).build();
                Log.e("Step 4",req.toString());
                Log.e("Is Shutdown", String.valueOf(mChannel.isShutdown()));
                Login.LoginResponse resp = stub.login(req);
                Log.e("Step 5",resp.toString());
                return resp.getLoginResponseData().getSuccess();
            } catch (Exception e) {
                Log.e("Is Shutdown", String.valueOf(mChannel.isShutdown()));
                Log.e("Some Error occured",e.toString());
                return "Error";
            }
        }

        @Override
        protected void onPostExecute(String result) {
            try {
                mChannel.shutdown().awaitTermination(1, TimeUnit.SECONDS);
            } catch (InterruptedException e) {
               Log.d("Channel:", String.valueOf(mChannel.isShutdown()));
                e.printStackTrace();
            }
        }
    }


}
