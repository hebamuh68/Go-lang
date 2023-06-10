from matplotlib import pyplot
import threading
import os

class AnalyzeSpectrum(object):
    def __init__(self):
        pass


    def Get_Spectrum_Indices_At_Level(self, level, msexperiment_object):
        indices = []
        for i in range(msexperiment_object.getNrSpectra()):
            if msexperiment_object[i].getMSLevel() == level:
                indices.append(i)

        return indices 



    def Get_MSData_At_Level(self, level, msexperiment_object):
        mz = []
        intensity = []
        retention_time = []

        for spectrum in msexperiment_object.getSpectra():
            if spectrum.getMSLevel() == level:
                mz.append(spectrum.get_peaks()[0])
                intensity.append(spectrum.get_peaks()[1])
                retention_time.append(spectrum.getRT())
        
        return mz, intensity, retention_time
    
    def Get_MSData(self, msexperiment_object, spectrum_indices):
        mz = []
        intensity = []
        retention_time = []
        
        for index in spectrum_indices:
            spectrum = msexperiment_object.getSpectrum(index)
            mz.append(spectrum.get_peaks()[0])
            intensity.append(spectrum.get_peaks()[1])
            retention_time.append(spectrum.getRT())
        
        return mz, intensity, retention_time
        
        
    def Draw_One_3D_Plot(self, axis, x_values, y_values, z_values):
        axis.plot(x_values, y_values, z_values)

    def Draw_3D_Spectrum(self, mz_data, intensity_data, retention_times, range_of_indices=None,
                         folder_path=None, sample='0'):
        pyplot.style.use("seaborn")
        fig = pyplot.figure(figsize=(20, 20))
        ax = fig.add_subplot(111, projection='3d')
        if range_of_indices == None:
            
            if len(mz_data) < 5 :
                for i in range(len(mz_data)):
                    self.Draw_One_3D_Plot(ax, mz_data[i], [retention_times[i]]*len(mz_data[i]), intensity_data[i])
                    
                ax.set_title(f"Sample: {sample}", fontdict={'fontsize':23, 'fontweight':'bold', 'color':'#00005C'})
                ax.set_xlabel('M/Z', fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#3B185F'})
                ax.set_ylabel("Retention Time", fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#C060A1'})
                ax.set_zlabel("Intensity", fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#F0CAA3'})
                pyplot.show();
                
            for i in range(0, len(mz_data), 5):
                if len(mz_data) - i < 5 :
                    for j in range(len(mz_data) - i):
                        self.Draw_One_3D_Plot(ax, mz_data[i], [retention_times[i]]*len(mz_data[i]), intensity_data[i])
                else:    
                    t1 = threading.Thread(target=self.Draw_One_3D_Plot, args=[ax, mz_data[i], [retention_times[i]]*len(mz_data[i]), intensity_data[i]])
                    t2 = threading.Thread(target=self.Draw_One_3D_Plot, args=[ax, mz_data[i+1], [retention_times[i+1]]*len(mz_data[i+1]), intensity_data[i+1]])
                    t3 = threading.Thread(target=self.Draw_One_3D_Plot, args=[ax, mz_data[i+2], [retention_times[i+2]]*len(mz_data[i+2]), intensity_data[i+2]])
                    t4 = threading.Thread(target=self.Draw_One_3D_Plot, args=[ax, mz_data[i+3], [retention_times[i+3]]*len(mz_data[i+3]), intensity_data[i+3]])
                    t5 = threading.Thread(target=self.Draw_One_3D_Plot, args=[ax, mz_data[i+4], [retention_times[i+4]]*len(mz_data[i+4]), intensity_data[i+4]])
                    t1.start()
                    t2.start()
                    t3.start()
                    t4.start()
                    t5.start()
                    t1.join()
                    t2.join()
                    t3.join()
                    t4.join()
                    t5.join()
                    
            ax.set_title(f"Sample: {sample}", fontdict={'fontsize':23, 'fontweight':'bold', 'color':'#00005C'})
            
        
        else:
            
            if range_of_indices[1] < 5 :
                for i in range(range_of_indices[0], range_of_indices[1]):
                    self.Draw_One_3D_Plot(ax, mz_data[i], [retention_times[i]]*len(mz_data[i]), intensity_data[i])
                    
                ax.set_title(f"Sample: {sample}", fontdict={'fontsize':23, 'fontweight':'bold', 'color':'#00005C'})
                ax.set_xlabel('M/Z', fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#3B185F'})
                ax.set_ylabel("Retention Time", fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#C060A1'})
                ax.set_zlabel("Intensity", fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#F0CAA3'})
                pyplot.show();
                
            for i in range(range_of_indices[0], range_of_indices[1], 5):
                if range_of_indices[1] - i < 5:
                    for i in range(len(mz_data)):
                        self.Draw_One_3D_Plot(ax, mz_data[i], [retention_times[i]]*len(mz_data[i]), intensity_data[i])

                else:   
                    t1 = threading.Thread(target=self.Draw_One_3D_Plot, args=[ax, mz_data[i], [retention_times[i]]*len(mz_data[i]), intensity_data[i]])
                    t2 = threading.Thread(target=self.Draw_One_3D_Plot, args=[ax, mz_data[i+1], [retention_times[i+1]]*len(mz_data[i+1]), intensity_data[i+1]])
                    t3 = threading.Thread(target=self.Draw_One_3D_Plot, args=[ax, mz_data[i+2], [retention_times[i+2]]*len(mz_data[i+2]), intensity_data[i+2]])
                    t4 = threading.Thread(target=self.Draw_One_3D_Plot, args=[ax, mz_data[i+3], [retention_times[i+3]]*len(mz_data[i+3]), intensity_data[i+3]])
                    t5 = threading.Thread(target=self.Draw_One_3D_Plot, args=[ax, mz_data[i+4], [retention_times[i+4]]*len(mz_data[i+4]), intensity_data[i+4]])
                    t1.start()
                    t2.start()
                    t3.start()
                    t4.start()
                    t5.start()
                    t1.join()
                    t2.join()
                    t3.join()
                    t4.join()
                    t5.join()
            ax.set_title(f"Spectra from {range_of_indices[0]} to {range_of_indices[1]}", fontdict={'fontsize':23, 'fontweight':'bold', 'color':'#00005C'})
            
            
        ax.set_xlabel('M/Z', fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#3B185F'})
        ax.set_ylabel("Retention Time", fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#C060A1'})
        ax.set_zlabel("Intensity", fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#F0CAA3'})
        if folder_path != None:
            fig_name = f'smaple_+{sample}.png'
            fig_path = os.path.join(folder_path, fig_name)
            pyplot.savefig(fig_path)
        else:
            pyplot.show();
        
        
    def Draw_Spectrum(self, mz_values, intensity_values, spectrum_index=None, color='#C060A1'):
        pyplot.style.use("seaborn")
        fig = pyplot.figure(figsize=(12, 8))
        ax = fig.add_subplot()
        if spectrum_index == None:
            ax.bar(mz_values, intensity_values, width=0.5, edgecolor=color, linewidth=0.9)
            ax.set(xlim=(0, max(mz_values)), ylim=(0, max(intensity_values)))
            ax.set_title("Spectrum",  fontdict={'fontsize':23, 'fontweight':'bold', 'color':'#00005C'})
        else:
            ax.bar(mz_values[spectrum_index], intensity_values[spectrum_index], width=0.5, edgecolor='#C060A1', linewidth=0.9)
            ax.set(xlim=(0, max(mz_values[spectrum_index])), ylim=(0, max(intensity_values[spectrum_index])))
            ax.set_title(f"Spectrum at index {spectrum_index}",  fontdict={'fontsize':23, 'fontweight':'bold', 'color':'#00005C'})

        ax.set_xlabel('M/Z', fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#3B185F'})
        ax.set_ylabel("Intensity", fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#F0CAA3'})
        pyplot.show();
    
    def Draw_Overlapped_Spectrums(self, mz_values1, intensity_values1, mz_values2, intensity_values2, color1='#C060A1', color2='navy'):
        pyplot.style.use("seaborn")
        fig = pyplot.figure(figsize=(12, 8))
        ax = fig.add_subplot()

        ax.bar(mz_values1, intensity_values1, width=0.5, edgecolor=color1, linewidth=0.9)
        ax.bar(mz_values2, intensity_values2, width=0.5, edgecolor=color2, linewidth=0.9)

        ax.set_title("Ovelapped Spectrums",  fontdict={'fontsize':23, 'fontweight':'bold', 'color':'#00005C'})
        ax.set_xlabel('M/Z', fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#3B185F'})
        ax.set_ylabel("Intensity", fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#F0CAA3'})
        pyplot.show();
    
    def Draw_Intensity_Over_Rt(self, mz_value, mz_data, intensity_data, retention_times, error=20):
        pyplot.style.use("seaborn")

        mass_error = (mz_value * error) / float(1e6)
        lower_interval = mz_value - mass_error
        upper_interval = mz_value + mass_error
        x = []
        y = []

        for i in range(len(mz_data)):
            for j in range(len(mz_data[i])):
                if mz_data[i][j] <= upper_interval and mz_data[i][j] > lower_interval:
                    x.append(retention_times[i])
                    y.append(intensity_data[i][j])
        
        fig = pyplot.figure(figsize=(10, 10))
        ax = fig.add_subplot()
        ax.plot(x, y)
        ax.set(xlim=(0, max(x)), ylim=(0, max(y)))
        ax.set_title('When M/Z ~= ' + str(mz_value), fontdict={'fontsize':23, 'fontweight':'bold', 'color':'#00005C'})
        ax.xlabel("Retention Time", fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#3B185F'})
        ax.ylabel("Intensity", fontdict={'fontsize':15, 'fontweight':'bold', 'color':'#F0CAA3'})
        pyplot.show();